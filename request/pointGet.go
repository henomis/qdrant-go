package request

import (
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type PointGet struct {
	CollectionName string  `json:"-"`
	ID             string  `json:"-"`
	Consistency    *string `json:"-"`
}

func (p *PointGet) Path() (string, error) {
	var urlValues restclientgo.URLValues
	urlValues.Add("consistency", p.Consistency)

	parameters := ""
	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s/points/%s%s", p.CollectionName, p.ID, parameters), nil
}

func (p *PointGet) Encode() (io.Reader, error) {
	return nil, nil
}

func (p *PointGet) ContentType() string {
	return ""
}
