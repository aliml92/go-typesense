package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentsService_MultiSearch(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/multi_search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
		{
			"results": [
				{
					"facet_counts": [],
					"found": 1,
					"hits": [
						{
							"document": {
								"brand": "Fila",
								"id": "1",
								"name": "shoe",
								"price": 87
							},
							"highlight": {
								"name": {
									"matched_tokens": [
										"shoe"
									],
									"snippet": "<mark>shoe</mark>"
								}
							},
							"highlights": [
								{
									"field": "name",
									"matched_tokens": [
										"shoe"
									],
									"snippet": "<mark>shoe</mark>"
								}
							],
							"text_match": 578730123365711993,
							"text_match_info": {
								"best_field_score": "1108091339008",
								"best_field_weight": 15,
								"fields_matched": 1,
								"score": "578730123365711993",
								"tokens_matched": 1
							}
						}
					],
					"out_of": 3,
					"page": 1,
					"request_params": {
						"collection_name": "products",
						"per_page": 10,
						"q": "shoe"
					},
					"search_cutoff": false,
					"search_time_ms": 7
				},
				{
					"code": 404,
					"error": "Not found."
				}
			]
		}
		`)
	})

	ctx := context.Background()
	body := &MultiSearchSearchesParameter{
		Searches: []MultiSearchCollectionParameters{
			{
				Collection: "products",
				Q:          String("shoe"),
				FilterBy:   String("price:=[50..120]"),
			},
			{
				Collection: "brandz", // let's say `brandz` collection do not exist
				Q:          String("Nike"),
			},
		},
	}
	opts := &MultiSearchParameters{
		QueryBy: String("name"),
	}
	got, err := client.Documents.MultiSearch(ctx, body, opts)
	assert.NoError(t, err)

	want := &MultiSearchResult{
		Results: []ResultOrError{
			{
				SearchResult: &SearchResult{
					FacetCounts: []*FacetCounts{},
					Found:       Int(1),
					Hits: []*SearchResultHit{{
						Document: map[string]interface{}{
							"brand": "Fila",
							"id":    "1",
							"name":  "shoe",
							"price": float64(87),
						},
						Highlight: map[string]interface{}{
							"name": map[string]interface{}{
								"matched_tokens": []interface{}{"shoe"},
								"snippet":        "<mark>shoe</mark>",
							},
						},
						Highlights: []*SearchHighlight{{
							Field:         String("name"),
							MatchedTokens: []string{"shoe"},
							Snippet:       String("<mark>shoe</mark>"),
						}},
						TextMatch: Int64(578730123365711993),
						TextMatchInfo: struct {
							BestFieldScore  string `json:"best_field_score"`
							BestFieldWeight int    `json:"best_field_weight"`
							FieldsMatched   int    `json:"fields_matched"`
							Score           string `json:"score"`
							TokensMatched   int    `json:"tokens_matched"`
						}{
							BestFieldScore:  "1108091339008",
							BestFieldWeight: 15,
							FieldsMatched:   1,
							Score:           "578730123365711993",
							TokensMatched:   1,
						},
					}},
					OutOf: Int(3),
					Page:  Int(1),
					RequestParams: &struct {
						CollectionName string `json:"collection_name"`
						PerPage        int    `json:"per_page"`
						Q              string `json:"q"`
					}{
						CollectionName: "products",
						PerPage:        10,
						Q:              "shoe",
					},
					SearchCutoff: Bool(false),
					SearchTimeMs: Int(7),
				},
			},
			{
				SearchError: &SearchError{
					Code:  404,
					Error: "Not found.",
				},
			},
		},
	}

	assert.Equal(t, want, got)
}
