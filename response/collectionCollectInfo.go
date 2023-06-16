package response

import (
	"encoding/json"
	"io"
)

type CollectionCollectInfo struct {
	Response
	Result CollectionCollectInfoResult `json:"result"`
}

type CollectionCollectInfoResult struct {
	Status          string           `json:"status"`
	OptimizerStatus string           `json:"optimizer_status"`
	VectorsCount    uint             `json:"vectors_count"`
	PointsCount     uint             `json:"points_count"`
	SegmentsCount   uint             `json:"segments_count"`
	Config          CollectionConfig `json:"config"`
}

type CollectionConfig struct {
	Params             VectorsParams    `json:"params"`
	HNSWConfig         HNSWConfig       `json:"hnsw_config"`
	OptimizerConfig    OptimizersConfig `json:"optimizer_config"`
	WalConfig          WALConfig        `json:"wal_config"`
	QuantizationConfig interface{}      `json:"quantization_config"`
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

func (l *CollectionCollectInfo) AcceptContentType() string {
	return "application/json"
}

func (l *CollectionCollectInfo) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(l)
}
