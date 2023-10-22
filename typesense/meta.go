package typesense

import (
	"context"
	"strings"
)

type MetaService service

// HealthStatus defines model for HealthStatus.
type HealthStatus struct {
	Ok bool `json:"ok"`
}

type NodeStatus struct {
	CommittedIndex int    `json:"committed_index"`
	QueuedWrites   int    `json:"queued_writes"`
	State          string `json:"state"`
}
type Config struct {
	LogSlowRequestsTimeMS *int  `json:"log-slow-requests-time-ms,omitempty"`
	HealthyReadLag        *int  `json:"healthy-read-lag,omitempty"`
	HealthyWriteLag       *int  `json:"healthy-write-lag,omitempty"`
	SkipWrites            *bool `json:"skip-writes,omitempty"`
}

type Metrics struct {
	SystemCPUActivePercentage         string            `json:"system_cpu_active_percentage"`
	SystemCPUIndividualPercentage     map[string]string `json:"-"`
	SystemDiskTotalBytes              string            `json:"system_disk_total_bytes"`
	SystemDiskUsedBytes               string            `json:"system_disk_used_bytes"`
	SystemMemoryTotalBytes            string            `json:"system_memory_total_bytes"`
	SystemMemoryUsedBytes             string            `json:"system_memory_used_bytes"`
	SystemNetworkReceivedBytes        string            `json:"system_network_received_bytes"`
	SystemNetworkSentBytes            string            `json:"system_network_sent_bytes"`
	TypesenseMemoryActiveBytes        string            `json:"typesense_memory_active_bytes"`
	TypesenseMemoryAllocatedBytes     string            `json:"typesense_memory_allocated_bytes"`
	TypesenseMemoryFragmentationRatio string            `json:"typesense_memory_fragmentation_ratio"`
	TypesenseMemoryMappedBytes        string            `json:"typesense_memory_mapped_bytes"`
	TypesenseMemoryMetadataBytes      string            `json:"typesense_memory_metadata_bytes"`
	TypesenseMemoryResidentBytes      string            `json:"typesense_memory_resident_bytes"`
	TypesenseMemoryRetainedBytes      string            `json:"typesense_memory_retained_bytes"`
}

type Stats struct {
	DeleteLatencyMS             float32            `json:"delete_latency_ms"`
	DeleteRequestsPerSecond     float32            `json:"delete_requests_per_second"`
	ImportLatencyMS             float32            `json:"import_latency_ms"`
	ImportRequestsPerSecond     float32            `json:"import_requests_per_second"`
	LatencyMS                   map[string]float32 `json:"latency_ms"`
	OverloadedRequestsPerSecond float32            `json:"overloaded_requests_per_second"`
	PendingWriteBatches         float32            `json:"pending_write_batches"`
	RequestsPerSecond           map[string]float32 `json:"requests_per_second"`
	SearchLatencyMS             float32            `json:"search_latency_ms"`
	SearchRequestsPerSecond     float32            `json:"search_requests_per_second"`
	TotalRequestsPerSecond      float32            `json:"total_requests_per_second"`
	WriteLatencyMS              float32            `json:"write_latency_ms"`
	WriteRequestsPerSecond      float32            `json:"write_requests_per_second"`
}

func (s *MetaService) Config(ctx context.Context, body *Config) (*SuccessStatus, error) {
	u := "/config"
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MetaService) Metrics(ctx context.Context) (*Metrics, error) {
	u := "/metrics.json"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res map[string]string
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	m := &Metrics{}
	m.SystemCPUIndividualPercentage = make(map[string]string)
	for key, value := range res {
		switch key {
		case "system_cpu_active_percentage":
			m.SystemCPUActivePercentage = value
		case "system_disk_total_bytes":
			m.SystemDiskTotalBytes = value
		case "system_disk_used_bytes":
			m.SystemDiskUsedBytes = value
		case "system_memory_total_bytes":
			m.SystemMemoryTotalBytes = value
		case "system_memory_used_bytes":
			m.SystemMemoryUsedBytes = value
		case "system_network_received_bytes":
			m.SystemNetworkReceivedBytes = value
		case "system_network_sent_bytes":
			m.SystemNetworkSentBytes = value
		case "typesense_memory_active_bytes":
			m.TypesenseMemoryActiveBytes = value
		case "typesense_memory_allocated_bytes":
			m.TypesenseMemoryAllocatedBytes = value
		case "typesense_memory_fragmentation_ratio":
			m.TypesenseMemoryFragmentationRatio = value
		case "typesense_memory_mapped_bytes":
			m.TypesenseMemoryMappedBytes = value
		case "typesense_memory_metadata_bytes":
			m.TypesenseMemoryMetadataBytes = value
		case "typesense_memory_resident_bytes":
			m.TypesenseMemoryResidentBytes = value
		case "typesense_memory_retained_bytes":
			m.TypesenseMemoryRetainedBytes = value
		default:
			if strings.HasPrefix(key, "system_cpu") && strings.HasSuffix(key, "active_percentage") {
				m.SystemCPUIndividualPercentage[key] = value
			}
		}
	}
	return m, nil
}

func (s *MetaService) Stats(ctx context.Context) (*Stats, error) {
	u := "/stats.json"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Stats{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Debug struct {
	State   int    `json:"state"`
	Version string `json:"version"`
}

func (s *MetaService) Debug(ctx context.Context) (*Debug, error) {
	u := "/debug"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &Debug{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MetaService) Health(ctx context.Context) (*HealthStatus, error) {
	u := "/health"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &HealthStatus{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MetaService) Status(ctx context.Context) (*NodeStatus, error) {
	u := "/status"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &NodeStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
