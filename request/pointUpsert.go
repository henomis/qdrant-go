package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type Ordering string

const (
	OrderingWeak   Ordering = "weak"
	OrderingMedium Ordering = "medium"
	OrderingStrong Ordering = "strong"
)

type PointUpsert struct {
	CollectionName string    `json:"-"`
	Wait           *bool     `json:"-"`
	Ordering       *Ordering `json:"-"`
	Points         []Point   `json:"points"`
}

type Point struct {
	ID      string                 `json:"id"`
	Vector  []float64              `json:"vector"`
	Payload map[string]interface{} `json:"payload,omitempty"`
}

func (p *PointUpsert) Path() (string, error) {
	path := fmt.Sprintf("/collections/%s/points", p.CollectionName)

	urlValues := restclientgo.NewURLValues()
	urlValues.AddBool("wait", p.Wait)
	urlValues.Add("ordering", (*string)(p.Ordering))

	urlValuesEncoded := urlValues.Encode()
	if urlValuesEncoded != "" {
		path = fmt.Sprintf("%s?%s", path, urlValuesEncoded)
	}

	return path, nil
}

func (p *PointUpsert) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (p *PointUpsert) ContentType() string {
	return "application/json"
}
