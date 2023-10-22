package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimitService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/limits", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			[ 
				{
					"id": 1,
					"action": "throttle",
					"api_keys": [
						"abc"
					],
					"auto_ban_1m_duration_hours": 1,
					"auto_ban_1m_threshold": 1,
					"ip_addresses": [
						"172.0.0.1"
					],
					"max_requests": {
						"hour_threshold": 10,
						"minute_threshold": 10
					},
					"priority": 0
				}
			]`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.List(ctx)
	assert.NoError(t, err)

	want := []*RateLimitRule{{
		ID:                     1,
		Action:                 "throttle",
		ApiKeys:                []string{"abc"},
		IpAddresses:           []string{"172.0.0.1"},
		AutoBan1mThreshold:     Int(1),
		AutoBan1mDurationHours: Int(1),
		MaxRequests: &struct {
			MinuteThreshold *int `json:"minute_threshold,omitempty"`
			HourThreshold   *int `json:"hour_threshold,omitempty"`
		}{
			MinuteThreshold: Int(10),
			HourThreshold:   Int(10),
		},
		Priority: Int(0),
	}}

	assert.Equal(t, want, got)
}

func TestRateLimitService_ListActive(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/limits/active", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			[ 
				{
					"status_id": 1,
					"throttling_from": 1000,
					"throttling_to": 2000,
					"ip_address": "172.0.0.1",
					"api_key": "abc"
				}
			]`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.ListActive(ctx)
	assert.NoError(t, err)

	want := []*RateLimitStatus{{
		StatusId:       1,
		ThrottlingFrom: 1000,
		ThrottlingTo:   2000,
		IpAddress:      String("172.0.0.1"),
		ApiKey:         String("abc"),
	}}

	assert.Equal(t, want, got)
}

func TestRateLimitService_ListExceeds(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/limits/exceeds", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			[ 
				{
					"id": 1,
					"api_key": "abc",
					"ip": "172.0.0.1",
					"request_count": 200
				}
			]`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.ListExceeds(ctx)
	assert.NoError(t, err)

	want := []*RateLimitExceed{{
		Id:           1,
		ApiKey:       "abc",
		Ip:           "172.0.0.1",
		RequestCount: 200,
	}}

	assert.Equal(t, want, got)
}

func TestRateLimitService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/limits", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"message": "Rule added successfully.",
				"rule": {
					"action": "throttle",
					"api_keys": [
						"abc"
					],
					"auto_ban_1m_duration_hours": 1,
					"auto_ban_1m_threshold": 1,
					"id": 1,
					"max_requests": {
						"minute_threshold": 10
					},
					"priority": 0
				}
			}
			`)
	})

	ctx := context.Background()
	body := &RateLimitRuleSchema{
		Action:                 "throttle",
		ApiKeys:                []string{"abc"},
		MaxRequests1m:          Int(10),
		MaxRequests1h:          Int(-1),
		AutoBan1mThreshold:     Int(1),
		AutoBan1mDurationHours: Int(1),
	}
	got, err := client.RateLimits.Create(ctx, body)
	assert.NoError(t, err)

	want := &RateLimitResponse{
		Message: "Rule added successfully.",
		Rule: &RateLimitRule{
			ID:                     1,
			Action:                 "throttle",
			ApiKeys:                []string{"abc"},
			AutoBan1mThreshold:     Int(1),
			AutoBan1mDurationHours: Int(1),
			MaxRequests: &struct {
				MinuteThreshold *int `json:"minute_threshold,omitempty"`
				HourThreshold   *int `json:"hour_threshold,omitempty"`
			}{
				MinuteThreshold: Int(10),
			},
			Priority: Int(0),
		},
	}

	assert.Equal(t, want, got)
}

func TestRateLimitService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	id := 1
	u := fmt.Sprintf("/limits/%d", id)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"action": "throttle",
				"api_keys": [
					"abc"
				],
				"auto_ban_1m_duration_hours": 1,
				"auto_ban_1m_threshold": 1,
				"id": 1,
				"max_requests": {
					"minute_threshold": 10
				},
				"priority": 0
			}
			`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.Get(ctx, id)
	assert.NoError(t, err)

	want := &RateLimitRule{
		ID:                     1,
		Action:                 "throttle",
		ApiKeys:                []string{"abc"},
		AutoBan1mThreshold:     Int(1),
		AutoBan1mDurationHours: Int(1),
		MaxRequests: &struct {
			MinuteThreshold *int `json:"minute_threshold,omitempty"`
			HourThreshold   *int `json:"hour_threshold,omitempty"`
		}{
			MinuteThreshold: Int(10),
		},
		Priority: Int(0),
	}

	assert.Equal(t, want, got)
}

func TestRateLimitService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	id := 1
	u := fmt.Sprintf("/limits/%d", id)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"id": 1
			}
		`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.Delete(ctx, id)
	assert.NoError(t, err)

	want := &DeleteRateLimitResponse{
		ID: 1,
	}

	assert.Equal(t, want, got)
}

func TestRateLimitService_DeleteActive(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	id := 1
	u := fmt.Sprintf("/limits/active/%d", id)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"id": 1
			}
		`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.DeleteActive(ctx, id)
	assert.NoError(t, err)

	want := &DeleteRateLimitResponse{
		ID: 1,
	}

	assert.Equal(t, want, got)
}

func TestRateLimitService_DeleteExceeds(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	id := 1
	u := fmt.Sprintf("/limits/active/%d", id)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"id": 1
			}
		`)
	})

	ctx := context.Background()
	got, err := client.RateLimits.DeleteActive(ctx, id)
	assert.NoError(t, err)

	want := &DeleteRateLimitResponse{
		ID: 1,
	}

	assert.Equal(t, want, got)
}
