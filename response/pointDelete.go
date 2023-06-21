package response

import (
	"encoding/json"
	"io"
)

type PointDelete struct {
	Response
	Result PointOperationResult `json:"result"`
}

func (p *PointDelete) AcceptContentType() string {
	return "application/json"
}

func (p *PointDelete) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(p)
}
