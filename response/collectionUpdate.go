package response

import (
	"encoding/json"
	"io"
)

type CollectionUpdate struct {
	Response
	Result bool `json:"result"`
}

func (c *CollectionUpdate) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionUpdate) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}

	c.Response.SetStatusMessage()
	return nil
}
