package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type CollectionCreate struct {
	CollectionName         string            `json:"-"`
	Vectors                VectorsParams     `json:"vectors"`
	ShardNumber            *int              `json:"shard_number,omitempty"`
	ReplicationFactor      *int              `json:"replication_factor,omitempty"`
	WriteConsistencyFactor *int              `json:"write_consistency_factor,omitempty"`
	OnDiskPayload          *bool             `json:"on_disk_payload,omitempty"`
	HNSWConfig             *HNSWConfig       `json:"hnsw_config,omitempty"`
	WALConfig              *WALConfig        `json:"wal_config,omitempty"`
	OptimizersConfig       *OptimizersConfig `json:"optimizers_config,omitempty"`
	InitFrom               interface{}       `json:"init_from,omitempty"`
	QuantizationConfig     interface{}       `json:"quantization_config,omitempty"`
}

type Distance string

const (
	DistanceCosine    Distance = "Cosine"
	DistanceEuclidean Distance = "Euclid"
	DistanceDot       Distance = "Dot"
)

type VectorsParams struct {
	Size               uint64      `json:"size"`
	Distance           Distance    `json:"distance"`
	HNSWConfig         *HNSWConfig `json:"hnsw_config,omitempty"`
	QuantizationConfig interface{} `json:"quantization_config,omitempty"`
	OnDisk             *bool       `json:"on_disk,omitempty"`
}

type HNSWConfig struct {
	M                  *uint `json:"m,omitempty"`
	EfConstruction     *uint `json:"ef_construction,omitempty"`
	FullScanThreshold  *uint `json:"full_scan_threshold,omitempty"`
	MaxIndexingThreads *uint `json:"max_indexing_threads,omitempty"`
	OnDisk             *bool `json:"on_disk,omitempty"`
	PayloadM           *uint `json:"payload_m,omitempty"`
}

type WALConfig struct {
	WalCapacityMB    *uint `json:"wal_capacity_mb,omitempty"`
	WalSegmentsAhead *uint `json:"wal_segments_ahead,omitempty"`
}

type OptimizersConfig struct {
	DeletedThreshold       *float32 `json:"deleted_threshold,omitempty"`
	VacuumMinVectorNumber  *uint    `json:"vacuum_min_vector_number,omitempty"`
	DefaultSegmentNumber   *uint    `json:"default_segment_number,omitempty"`
	MaxSegmentSize         *uint    `json:"max_segment_size,omitempty"`
	MemMapThreshold        *uint    `json:"memmap_threshold,omitempty"`
	IndexingThreshold      *uint    `json:"indexing_threshold,omitempty"`
	FlushIntervalSec       *uint64  `json:"flush_interval_sec,omitempty"`
	MaxOptimizationThreads *uint    `json:"max_optimization_threads,omitempty"`
}

type InitFrom struct {
	Collection string `json:"collection"`
}

func (c *CollectionCreate) Path() (string, error) {
	return fmt.Sprintf("/collections/%s", c.CollectionName), nil
}

func (c *CollectionCreate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *CollectionCreate) ContentType() string {
	return "application/json"
}
