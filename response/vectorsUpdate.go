package response

import (
	"encoding/json"
	"io"
)

type VectorsUpdate struct {
	Response
	Result OperationResult `json:"result"`
}

func (p *VectorsUpdate) AcceptContentType() string {
	return "application/json"
}

func (p *VectorsUpdate) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
