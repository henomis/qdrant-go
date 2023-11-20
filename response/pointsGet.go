package response

import (
	"encoding/json"
	"io"
)

type PointsGet struct {
	Response
	Result []PointsGetResult `json:"result"`
}

type PointsGetResult struct {
	ID      string         `json:"id"`
	Payload map[string]any `json:"payload"`
	Vector  []float64      `json:"vector"`
}

func (p *PointsGet) AcceptContentType() string {
	return "application/json"
}

func (p *PointsGet) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
