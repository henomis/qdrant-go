package response

import (
	"encoding/json"
	"io"
)

type PointUpsert struct {
	Response
	Result PointUpsertResult `json:"result"`
}

type PointUpsertResultStatus string

const (
	PointUpsertResultStatusAcknowledged PointUpsertResultStatus = "acknowledged"
	PointUpsertResultStatusCompleted    PointUpsertResultStatus = "completed"
)

type PointUpsertResult struct {
	OperationID uint64                  `json:"operation_id"`
	Status      PointUpsertResultStatus `json:"status"`
}

func (p *PointUpsert) AcceptContentType() string {
	return "application/json"
}

func (p *PointUpsert) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(p)
}
