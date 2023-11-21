package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type IndexCreate struct {
	CollectionName string    `json:"-"`
	Wait           *bool     `json:"-"`
	Ordering       *Ordering `json:"-"`
	FieldName      string    `json:"field_name"`
	FieldSchema    any       `json:"field_schema,omitempty"`
}

func (c *IndexCreate) Path() (string, error) {
	var urlValues = restclientgo.URLValues{}
	urlValues.AddBool("wait", c.Wait)
	urlValues.Add("ordering", (*string)(c.Ordering))

	parameters := ""
	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s/index%s", c.CollectionName, parameters), nil
}

func (c *IndexCreate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *IndexCreate) ContentType() string {
	return "application/json"
}
