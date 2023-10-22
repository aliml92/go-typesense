package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyticsRulesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/analytics/rules", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"rules": [
				{
					"name": "product_queries_aggregation",
					"params": {
						"destination": {
							"collection": "product_queries"
						},
						"limit": 1000,
						"source": {
							"collections": [
								"products"
							]
						}
					},
					"type": "popular_queries"
				}
			]
		}`)
	})

	ctx := context.Background()
	got, err := client.AnalyticsRules.List(ctx)
	assert.NoError(t, err)

	want := &AnalyticsRuleListResponse{
		Rules: []*AnalyticsRule{{
			Name: "product_queries_aggregation",
			Params: AnalyticsRuleParams{
				Source: struct {
					Collections []string `json:"collections"`
				}{
					Collections: []string{"products"},
				},
				Destination: struct {
					Collection string `json:"collection"`
				}{
					Collection: "product_queries",
				},
				Limit: 1000,
			},
			Type: "popular_queries",
		}},
	}
	assert.Equal(t, want, got)
}

func TestAnalyticsRulesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/analytics/rules", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "product_queries_aggregation",
			"params": {
				"destination": {
					"collection": "product_queries"
				},
				"limit": 1000,
				"source": {
					"collections": [
						"products"
					]
				}
			},
			"type": "popular_queries"
		}`)
	})

	ctx := context.Background()
	body := &AnalyticsRule{
		Name: "product_queries_aggregation",
		Params: AnalyticsRuleParams{
			Source: struct {
				Collections []string `json:"collections"`
			}{
				Collections: []string{"products"},
			},
			Destination: struct {
				Collection string `json:"collection"`
			}{
				Collection: "product_queries",
			},
			Limit: 1000,
		},
		Type: "popular_queries",
	}
	got, err := client.AnalyticsRules.Create(ctx, body)
	assert.NoError(t, err)

	assert.Equal(t, body, got)
}

func TestAnalyticsRulesService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	ruleName := "product_queries_aggregation"
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "product_queries_aggregation",
			"params": {
				"destination": {
					"collection": "product_queries"
				},
				"limit": 1000,
				"source": {
					"collections": [
						"products"
					]
				}
			},
			"type": "popular_queries"
		}`)
	})

	ctx := context.Background()

	got, err := client.AnalyticsRules.Get(ctx, ruleName)
	assert.NoError(t, err)
	want := &AnalyticsRule{
		Name: "product_queries_aggregation",
		Params: AnalyticsRuleParams{
			Source: struct {
				Collections []string `json:"collections"`
			}{
				Collections: []string{"products"},
			},
			Destination: struct {
				Collection string `json:"collection"`
			}{
				Collection: "product_queries",
			},
			Limit: 1000,
		},
		Type: "popular_queries",
	}

	assert.Equal(t, want, got)
}

func TestAnalyticsRulesService_Upsert(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	ruleName := "product_queries_aggregation"
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "product_queries_aggregation",
			"params": {
				"destination": {
					"collection": "product_queries"
				},
				"limit": 1000,
				"source": {
					"collections": [
						"products"
					]
				}
			},
			"type": "popular_queries"
		}`)
	})

	ctx := context.Background()
	body := &AnalyticsRuleUpsertSchema{
		Params: AnalyticsRuleParams{
			Source: struct {
				Collections []string `json:"collections"`
			}{
				Collections: []string{"products"},
			},
			Destination: struct {
				Collection string `json:"collection"`
			}{
				Collection: "product_queries",
			},
			Limit: 1000,
		},
		Type: "popular_queries",
	}
	got, err := client.AnalyticsRules.Upsert(ctx, ruleName, body)
	assert.NoError(t, err)
	want := &AnalyticsRule{
		Name: "product_queries_aggregation",
		Params: AnalyticsRuleParams{
			Source: struct {
				Collections []string `json:"collections"`
			}{
				Collections: []string{"products"},
			},
			Destination: struct {
				Collection string `json:"collection"`
			}{
				Collection: "product_queries",
			},
			Limit: 1000,
		},
		Type: "popular_queries",
	}

	assert.Equal(t, want, got)
}

func TestAnalyticsRulesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	ruleName := "product_queries_aggregation"
	u := fmt.Sprintf("/analytics/rules/%s", ruleName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "product_queries_aggregation"
		}`)
	})

	ctx := context.Background()

	got, err := client.AnalyticsRules.Delete(ctx, ruleName)
	assert.NoError(t, err)
	want := &AnalyticsRuleDeleteResponse{
		Name: "product_queries_aggregation",
	}

	assert.Equal(t, want, got)
}
