package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type VectorsDelete struct {
	CollectionName string    `json:"-"`
	Points         []string  `json:"points,omitempty"`
	Filter         Filter    `json:"filter,omitempty"`
	Vector         []string  `json:"vector"`
	Wait           *bool     `json:"-"`
	Ordering       *Ordering `json:"-"`
}

func (p *VectorsDelete) Path() (string, error) {
	path := fmt.Sprintf("/collections/%s/points/vectors/delete", p.CollectionName)

	urlValues := restclientgo.NewURLValues()
	urlValues.Add("ordering", (*string)(p.Ordering))
	urlValues.AddBool("wait", p.Wait)

	urlValuesEncoded := urlValues.Encode()
	if urlValuesEncoded != "" {
		path = fmt.Sprintf("%s?%s", path, urlValuesEncoded)
	}

	return path, nil
}

func (p *VectorsDelete) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (p *VectorsDelete) ContentType() string {
	return "application/json"
}
