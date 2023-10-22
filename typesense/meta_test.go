package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetaService_Config(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"success": true
			}
		`)
	})

	ctx := context.Background()
	body := &Config{
		LogSlowRequestsTimeMS: Int(1000),
		HealthyReadLag:        Int(1000),
		HealthyWriteLag:       Int(1000),
		SkipWrites:            Bool(false),
	}
	got, err := client.Meta.Config(ctx, body)
	assert.NoError(t, err)
	want := &SuccessStatus{
		Success: true,
	}

	assert.Equal(t, want, got)
}

func TestMetaService_Metrics(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/metrics.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"system_cpu10_active_percentage": "0.00",
			"system_cpu11_active_percentage": "0.00",
			"system_cpu12_active_percentage": "20.00",
			"system_cpu1_active_percentage": "30.77",
			"system_cpu2_active_percentage": "11.11",
			"system_cpu3_active_percentage": "44.44",
			"system_cpu4_active_percentage": "0.00",
			"system_cpu5_active_percentage": "10.00",
			"system_cpu6_active_percentage": "0.00",
			"system_cpu7_active_percentage": "0.00",
			"system_cpu8_active_percentage": "10.00",
			"system_cpu9_active_percentage": "11.11",
			"system_cpu_active_percentage": "9.24",
			"system_disk_total_bytes": "241369505792",
			"system_disk_used_bytes": "116254326784",
			"system_memory_total_bytes": "7618924544",
			"system_memory_used_bytes": "6258020352",
			"system_network_received_bytes": "263031",
			"system_network_sent_bytes": "5210",
			"typesense_memory_active_bytes": "45043712",
			"typesense_memory_allocated_bytes": "39953112",
			"typesense_memory_fragmentation_ratio": "0.11",
			"typesense_memory_mapped_bytes": "166768640",
			"typesense_memory_metadata_bytes": "16807696",
			"typesense_memory_resident_bytes": "45043712",
			"typesense_memory_retained_bytes": "84889600"
		}
		`)
	})

	ctx := context.Background()
	got, err := client.Meta.Metrics(ctx)
	assert.NoError(t, err)
	want := &Metrics{
		SystemCPUActivePercentage: "9.24",
		SystemCPUIndividualPercentage: map[string]string{
			"system_cpu10_active_percentage": "0.00",
			"system_cpu11_active_percentage": "0.00",
			"system_cpu12_active_percentage": "20.00",
			"system_cpu1_active_percentage":  "30.77",
			"system_cpu2_active_percentage":  "11.11",
			"system_cpu3_active_percentage":  "44.44",
			"system_cpu4_active_percentage":  "0.00",
			"system_cpu5_active_percentage":  "10.00",
			"system_cpu6_active_percentage":  "0.00",
			"system_cpu7_active_percentage":  "0.00",
			"system_cpu8_active_percentage":  "10.00",
			"system_cpu9_active_percentage":  "11.11",
		},
		SystemDiskTotalBytes:              "241369505792",
		SystemDiskUsedBytes:               "116254326784",
		SystemMemoryTotalBytes:            "7618924544",
		SystemMemoryUsedBytes:             "6258020352",
		SystemNetworkReceivedBytes:        "263031",
		SystemNetworkSentBytes:            "5210",
		TypesenseMemoryActiveBytes:        "45043712",
		TypesenseMemoryAllocatedBytes:     "39953112",
		TypesenseMemoryFragmentationRatio: "0.11",
		TypesenseMemoryMappedBytes:        "166768640",
		TypesenseMemoryMetadataBytes:      "16807696",
		TypesenseMemoryResidentBytes:      "45043712",
		TypesenseMemoryRetainedBytes:      "84889600",
	}

	assert.Equal(t, want, got)
}

func TestMetaService_Stats(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stats.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"delete_latency_ms": 0,
			"delete_requests_per_second": 0,
			"import_latency_ms": 0,
			"import_requests_per_second": 0,
			"latency_ms": {
				"GET /collections/companies/documents/search": 0.0
			},
			"overloaded_requests_per_second": 0,
			"pending_write_batches": 0,
			"requests_per_second": {
				"GET /collections/companies/documents/search": 1.5
			},
			"search_latency_ms": 0.0,
			"search_requests_per_second": 1.5,
			"total_requests_per_second": 1.5,
			"write_latency_ms": 0,
			"write_requests_per_second": 0
		}
		`)
	})

	ctx := context.Background()
	got, err := client.Meta.Stats(ctx)
	assert.NoError(t, err)
	want := &Stats{
		DeleteLatencyMS:         0,
		DeleteRequestsPerSecond: 0,
		ImportLatencyMS:         0,
		ImportRequestsPerSecond: 0,
		LatencyMS: map[string]float32{
			"GET /collections/companies/documents/search": 0,
		},
		OverloadedRequestsPerSecond: 0,
		PendingWriteBatches:         0,
		RequestsPerSecond: map[string]float32{
			"GET /collections/companies/documents/search": 1.5,
		},
		SearchLatencyMS:         0,
		SearchRequestsPerSecond: 1.5,
		TotalRequestsPerSecond:  1.5,
		WriteLatencyMS:          0,
		WriteRequestsPerSecond:  0,
	}
	assert.Equal(t, want, got)
}

func TestMetaService_Debug(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"state": 1,
			"version": "0.25.0"
		}
		`)
	})

	ctx := context.Background()
	got, err := client.Meta.Debug(ctx)
	assert.NoError(t, err)

	want := &Debug{
		State:   1,
		Version: "0.25.0",
	}

	assert.Equal(t, want, got)
}

func TestMetaService_Health(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `
		{
			"ok": true
		}
		`)
	})

	ctx := context.Background()
	got, err := client.Meta.Health(ctx)
	assert.NoError(t, err)

	want := &HealthStatus{
		Ok: true,
	}

	assert.Equal(t, want, got)
}

func TestMetaService_Status(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"committed_index": 250,
			"queued_writes": 0,
			"state": "LEADER"
		}
		`)
	})

	ctx := context.Background()
	got, err := client.Meta.Status(ctx)
	assert.NoError(t, err)

	want := &NodeStatus{
		CommittedIndex: 250,
		QueuedWrites:   0,
		State:          "LEADER",
	}

	assert.Equal(t, want, got)
}
