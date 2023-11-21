package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type PointVectors struct {
	ID     string    `json:"id"`
	Vector []float64 `json:"vector"`
}

type VectorsUpdate struct {
	CollectionName string         `json:"-"`
	Points         []PointVectors `json:"points"`
	Wait           *bool          `json:"-"`
	Ordering       *Ordering      `json:"-"`
}

func (p *VectorsUpdate) Path() (string, error) {
	path := fmt.Sprintf("/collections/%s/points/vectors", p.CollectionName)

	urlValues := restclientgo.NewURLValues()
	urlValues.Add("ordering", (*string)(p.Ordering))
	urlValues.AddBool("wait", p.Wait)

	urlValuesEncoded := urlValues.Encode()
	if urlValuesEncoded != "" {
		path = fmt.Sprintf("%s?%s", path, urlValuesEncoded)
	}

	return path, nil
}

func (p *VectorsUpdate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (p *VectorsUpdate) ContentType() string {
	return "application/json"
}
