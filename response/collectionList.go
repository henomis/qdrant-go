package response

import (
	"encoding/json"
	"io"
)

type CollectionList struct {
	Response
	Result CollectionListResult `json:"result"`
}

type CollectionListResult struct {
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name string `json:"name"`
}

func (c *CollectionList) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionList) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}

	c.Response.SetStatusMessage()
	return nil
}
