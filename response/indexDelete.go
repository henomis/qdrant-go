package response

import (
	"encoding/json"
	"io"
)

type IndexDelete struct {
	Response
	Result OperationResult `json:"result"`
}

func (c *IndexDelete) AcceptContentType() string {
	return "application/json"
}

func (c *IndexDelete) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}

	c.Response.SetStatusMessage()
	return nil
}
