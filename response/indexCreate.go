package response

import (
	"encoding/json"
	"io"
)

type IndexCreate struct {
	Response
	Result OperationResult `json:"result"`
}

func (c *IndexCreate) AcceptContentType() string {
	return "application/json"
}

func (c *IndexCreate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
