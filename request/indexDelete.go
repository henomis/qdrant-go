package request

import (
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type IndexDelete struct {
	CollectionName string    `json:"-"`
	FieldName      string    `json:"-"`
	Wait           *bool     `json:"-"`
	Ordering       *Ordering `json:"-"`
}

func (c *IndexDelete) Path() (string, error) {
	var urlValues = restclientgo.URLValues{}
	urlValues.AddBool("wait", c.Wait)
	urlValues.Add("ordering", (*string)(c.Ordering))

	parameters := ""
	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s/index/%s%s", c.CollectionName, c.FieldName, parameters), nil
}

func (c *IndexDelete) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *IndexDelete) ContentType() string {
	return ""
}
