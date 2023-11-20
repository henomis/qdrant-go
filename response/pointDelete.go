package response

import (
	"encoding/json"
	"io"
)

type PointDelete struct {
	Response
	Result OperationResult `json:"result"`
}

func (p *PointDelete) AcceptContentType() string {
	return "application/json"
}

func (p *PointDelete) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
