package request

import (
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type CollectionDelete struct {
	CollectionName string `json:"-"`
	Timeout        *int   `json:"-"`
}

func (c *CollectionDelete) Path() (string, error) {
	var urlValues = restclientgo.URLValues{}
	urlValues.AddInt("timeout", c.Timeout)

	parameters := ""
	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s%s", c.CollectionName, parameters), nil
}

func (c *CollectionDelete) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *CollectionDelete) ContentType() string {
	return ""
}
