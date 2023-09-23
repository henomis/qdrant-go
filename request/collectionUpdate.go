package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type CollectionUpdate struct {
	CollectionName     string               `json:"-"`
	Timeout            *int                 `json:"timeout,omitempty"`
	Vectors            *UpdateVectorsParams `json:"vectors,omitempty"`
	OptimizersConfig   *OptimizersConfig    `json:"optimizers_config,omitempty"`
	Params             *UpdateParams        `json:"params,omitempty"`
	HNSWConfig         *HNSWConfig          `json:"hnsw_config,omitempty"`
	QuantizationConfig any                  `json:"quantization_config,omitempty"`
}

type UpdateVectorsParams struct {
	HNSWConfig         *HNSWConfig `json:"hnsw_config,omitempty"`
	QuantizationConfig any         `json:"quantization_config,omitempty"`
	OnDisk             *bool       `json:"on_disk,omitempty"`
}

type UpdateParams struct {
	ReplicationFactor      *int  `json:"replication_factor,omitempty"`
	WriteConsistencyFactor *int  `json:"write_consistency_factor,omitempty"`
	OnDiskPayload          *bool `json:"on_disk_payload,omitempty"`
}

func (c *CollectionUpdate) Path() (string, error) {
	var urlValues = restclientgo.URLValues{}
	var parameters string
	urlValues.AddInt("timeout", c.Timeout)

	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s%s", c.CollectionName, parameters), nil
}

func (c *CollectionUpdate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *CollectionUpdate) ContentType() string {
	return "application/json"
}
