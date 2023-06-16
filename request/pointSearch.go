package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type PointSearch struct {
	CollectionName string             `json:"-"`
	Consistency    *string            `json:"-"`
	Vector         []float64          `json:"vector"`
	Filter         interface{}        `json:"filter,omitempty"`
	Params         *PointSearchParams `json:"params,omitempty"`
	Limit          int                `json:"limit"`
	Offset         int                `json:"offset"`
	WithPayload    *bool              `json:"with_payload"`
	WithVector     *bool              `json:"with_vector,omitempty"`
	ScoreThreshold *float64           `json:"score_threshold,omitempty"`
}

type PointSearchParams struct {
	HNSWEF       *int                           `json:"hnsw_ef,omitempty"`
	Exact        bool                           `json:"exact"`
	Quantization *PointSearchParamsQuantization `json:"quantization,omitempty"`
}

type PointSearchParamsQuantization struct {
	Ignore  bool `json:"ignore"`
	Rescore bool `json:"rescore"`
}

func (p *PointSearch) Path() (string, error) {
	path := fmt.Sprintf("/collections/%s/points/search", p.CollectionName)

	urlValues := restclientgo.NewURLValues()
	urlValues.Add("ordering", (*string)(p.Consistency))

	urlValuesEncoded := urlValues.Encode()
	if urlValuesEncoded != "" {
		path = fmt.Sprintf("%s?%s", path, urlValuesEncoded)
	}

	return path, nil
}

func (p *PointSearch) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (p *PointSearch) ContentType() string {
	return "application/json"
}
