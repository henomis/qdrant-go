package response

import (
	"encoding/json"
	"io"
)

type CollectionDelete struct {
	Response
	Result bool `json:"result"`
}

func (c *CollectionDelete) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionDelete) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}

	c.Response.SetStatusMessage()
	return nil
}
