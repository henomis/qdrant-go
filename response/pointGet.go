package response

import (
	"encoding/json"
	"io"
)

type PointGet struct {
	Response
	Result PointGetResult `json:"result"`
}

type PointGetResult struct {
	ID      string         `json:"id"`
	Payload map[string]any `json:"payload"`
	Vector  []float64      `json:"vector"`
}

func (p *PointGet) AcceptContentType() string {
	return "application/json"
}

func (p *PointGet) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
