package response

import (
	"encoding/json"
	"io"
)

type CollectionCreate struct {
	Response
	Result bool `json:"result"`
}

func (c *CollectionCreate) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionCreate) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}

	c.Response.SetStatusMessage()
	return nil
}
