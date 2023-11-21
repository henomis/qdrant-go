package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type PointsDelete struct {
	CollectionName string    `json:"-"`
	Ordering       *Ordering `json:"-"`
	Wait           *bool     `json:"-"`
	Filter         Filter    `json:"filter,omitempty"`
	Points         []string  `json:"points,omitempty"`
}

type M map[string]interface{}

type Filter struct {
	Should  []M `json:"should,omitempty"`
	Must    []M `json:"must,omitempty"`
	MustNot []M `json:"must_not,omitempty"`
}

func (p *PointsDelete) Path() (string, error) {
	path := fmt.Sprintf("/collections/%s/points/delete", p.CollectionName)

	urlValues := restclientgo.NewURLValues()
	urlValues.Add("ordering", (*string)(p.Ordering))
	urlValues.AddBool("wait", p.Wait)

	urlValuesEncoded := urlValues.Encode()
	if urlValuesEncoded != "" {
		path = fmt.Sprintf("%s?%s", path, urlValuesEncoded)
	}

	return path, nil
}

func (p *PointsDelete) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (p *PointsDelete) ContentType() string {
	return "application/json"
}
