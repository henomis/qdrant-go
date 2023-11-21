package response

import (
	"encoding/json"
	"io"
)

type PointsDelete struct {
	Response
	Result OperationResult `json:"result"`
}

func (p *PointsDelete) AcceptContentType() string {
	return "application/json"
}

func (p *PointsDelete) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
