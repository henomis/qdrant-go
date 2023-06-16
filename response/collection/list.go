package collection

import (
	"encoding/json"
	"io"

	"github.com/henomis/qdrant-go/response"
)

type List struct {
	response.Response
	Result []ListResult `json:"result"`
}

type ListResult struct {
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name string `json:"name"`
}

func (l *List) AcceptContentType() string {
	return "application/json"
}

func (l *List) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(l)
}
