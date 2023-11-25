package response

import (
	"encoding/json"
	"io"
)

type PointsUpsert struct {
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

func (p *PointsUpsert) AcceptContentType() string {
	return "application/json"
}

func (p *PointsUpsert) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
