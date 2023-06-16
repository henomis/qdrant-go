package response

import (
	"encoding/json"
	"io"
)

type PointSearch struct {
	Response
	Result []PointSearchResult `json:"result"`
}

type PointSearchResult struct {
	ID      string                 `json:"id"`
	Version uint64                 `json:"version"`
	Score   float64                `json:"score"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Vector  []float64              `json:"vector,omitempty"`
}

func (p *PointSearch) AcceptContentType() string {
	return "application/json"
}

func (p *PointSearch) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(p)
}
