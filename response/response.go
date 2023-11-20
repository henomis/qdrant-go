package response

import (
	"io"

	"github.com/henomis/restclientgo"
)

type Response struct {
	Code      int     `json:"-"`
	Time      float64 `json:"time"`
	RawStatus any     `json:"status"`
	Status    string  `json:"-"`
	RawBody   *string `json:"-"`
}

type Detail struct {
	TypeURL *string `json:"typeUrl,omitempty"`
	Value   *string `json:"value,omitempty"`
}

func (r *Response) IsSuccess() bool {
	return (r.Code >= 200 && r.Code < 300)
}

func (r *Response) SetStatusCode(code int) error {
	r.Code = code
	return nil
}

func (r *Response) SetBody(body io.Reader) error {

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	s := string(b)
	r.RawBody = &s

	return nil
}

func (r *Response) SetHeaders(headers restclientgo.Headers) error {
	return nil
}

func (r *Response) SetStatusMessage() {
	switch status := r.RawStatus.(type) {
	case string:
		r.Status = status
	case map[string]interface{}:
		if errorMessage, ok := status["error"].(string); ok {
			r.Status = errorMessage
		}
	}
}
