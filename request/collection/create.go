package collection

import (
	"fmt"
	"io"
)

type Create struct {
	Name                   string            `json:"-"`
	Vectors                VectorsParams     `json:"vectors"`
	ShardNumber            *int              `json:"shard_number,omitempty"`
	ReplicationFactor      *int              `json:"replication_factor,omitempty"`
	WriteConsistencyFactor *int              `json:"write_consistency_factor,omitempty"`
	OnDiskPayload          *bool             `json:"on_disk_payload,omitempty"`
	HNSWConfig             *HNSWConfig       `json:"hnsw_config,omitempty"`
	WALConfig              *WALConfig        `json:"wal_config,omitempty"`
	OptimizersConfig       *OptimizersConfig `json:"optimizers_config,omitempty"`
	InitFrom               *InitFrom         `json:"init_from,omitempty"`
}

type VectorsParams struct {
}

type HNSWConfig struct {
}

type WALConfig struct {
}

type OptimizersConfig struct {
}

type InitFrom struct {
	Collection string `json:"collection"`
}

func (r *Create) Path() (string, error) {
	return fmt.Sprintf("/collections/%s", r.Name), nil
}

func (l *Create) Encode() (io.Reader, error) {
	return nil, nil
}

func (l *Create) ContentType() string {
	return ""
}
