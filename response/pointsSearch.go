package response

import (
	"encoding/json"
	"io"
)

type PointsSearch struct {
	Response
	Result []PointsSearchResult `json:"result"`
}

type PointsSearchResult struct {
	ID      string                 `json:"id"`
	Version uint64                 `json:"version"`
	Score   float64                `json:"score"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Vector  []float64              `json:"vector,omitempty"`
}

func (p *PointsSearch) AcceptContentType() string {
	return "application/json"
}

func (p *PointsSearch) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
