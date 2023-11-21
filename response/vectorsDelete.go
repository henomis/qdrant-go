package response

import (
	"encoding/json"
	"io"
)

type VectorsDelete struct {
	Response
	Result OperationResult `json:"result"`
}

func (p *VectorsDelete) AcceptContentType() string {
	return "application/json"
}

func (p *VectorsDelete) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
