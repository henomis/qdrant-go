package response

import (
	"encoding/json"
	"io"
)

type PointUpsert struct {
	Response
	Result OperationResult `json:"result"`
}

type OperationResultStatus string

const (
	PointOperationResultStatusAcknowledged OperationResultStatus = "acknowledged"
	PointOperationResultStatusCompleted    OperationResultStatus = "completed"
)

type OperationResult struct {
	OperationID uint64                `json:"operation_id"`
	Status      OperationResultStatus `json:"status"`
}

func (p *PointUpsert) AcceptContentType() string {
	return "application/json"
}

func (p *PointUpsert) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(p)
}
