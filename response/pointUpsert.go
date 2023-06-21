package response

import (
	"encoding/json"
	"io"
)

type PointUpsert struct {
	Response
	Result PointOperationResult `json:"result"`
}

type PointOperationResultStatus string

const (
	PointOperationResultStatusAcknowledged PointOperationResultStatus = "acknowledged"
	PointOperationResultStatusCompleted    PointOperationResultStatus = "completed"
)

type PointOperationResult struct {
	OperationID uint64                     `json:"operation_id"`
	Status      PointOperationResultStatus `json:"status"`
}

func (p *PointUpsert) AcceptContentType() string {
	return "application/json"
}

func (p *PointUpsert) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(p)
}
