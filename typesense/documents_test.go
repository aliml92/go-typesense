package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"id": "124",
			"name": "shoe",
			"price": 87.5,
			"rating": 4
		}`)
	})

	want := map[string]interface{}{
		"id":     "124",
		"name":   "shoe",
		"price":  87.5,
		"rating": float64(4),
	}
	ctx := context.Background()
	got, err := client.Documents.Create(ctx, collectionName, want)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	documentId := "124"
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"id": "124",
			"company_name": "Stark Industries",
			"num_employees": 5215,
			"country": "USA"
		}`)
	})

	want := map[string]interface{}{
		"id":            "124",
		"company_name":  "Stark Industries",
		"num_employees": float64(5215),
		"country":       "USA",
	}
	ctx := context.Background()
	got, err := client.Documents.Get(ctx, collectionName, documentId)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	documentId := "124"
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"company_name": "Stark Industries",
			"num_employees": 5215
		}`)
	})

	want := map[string]interface{}{
		"company_name":  "Stark Industries",
		"num_employees": float64(5215),
	}
	ctx := context.Background()
	got, err := client.Documents.Get(ctx, collectionName, documentId)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_UpdateByQuery(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"

	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"num_updated": 4
		}`)
	})

	want := &UpdateByQueryResponse{
		NumUpdated: 4,
	}

	ctx := context.Background()
	body := struct {
		CompanyName  string `json:"company_name"`
		NumEmployees int    `json:"num_employees"`
	}{
		CompanyName:  "Stark Industries",
		NumEmployees: 5500,
	}

	opts := &UpdateOptions{
		FilterBy:    "num_employees:>1000",
		DirtyValues: CoerceOrDrop,
	}
	got, err := client.Documents.UpdateByQuery(ctx, collectionName, &body, opts)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	documentId := "124"
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"id": "124",
			"company_name": "Stark Industries",
			"num_employees": 5215,
			"country": "USA"
		}`)
	})

	want := map[string]interface{}{
		"id":            "124",
		"company_name":  "Stark Industries",
		"num_employees": float64(5215),
		"country":       "USA",
	}

	ctx := context.Background()

	got, err := client.Documents.Delete(ctx, collectionName, documentId)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_DeleteByQuery(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"num_deleted": 4
		}`)
	})

	want := &DeleteByQueryResponse{
		NumDeleted: Int(4),
	}

	ctx := context.Background()
	opts := &DeleteOptions{
		FilterBy:  "num_employees:>1000",
		BatchSize: 10,
	}
	got, err := client.Documents.DeleteByQuery(ctx, collectionName, opts)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Export(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents/export", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{"company_name": "Stark Industries", "num_employees": 5215}
			{"company_name": "Future Technology", "num_employees": 1232}
		`)
	})

	want := []map[string]interface{}{
		{
			"company_name": "Stark Industries", "num_employees": float64(5215),
		},
		{
			"company_name": "Future Technology", "num_employees": float64(1232),
		},
	}

	ctx := context.Background()
	opts := &ExportDocumentsParams{
		FilterBy: "num_employees:>1000",
	}
	got, err := client.Documents.Export(ctx, collectionName, opts)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Import(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents/import", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{"success": true}
			{"success": true}
		`)
	})

	want := []*ImportDocumentResponse{
		{Success: true},
		{Success: true},
	}

	ctx := context.Background()
	body := []map[string]interface{}{
		{"id": "1", "company_name": "Stark Industries", "num_employees": 5215, "country": "USA"},
		{"id": "2", "company_name": "Orbit Inc.", "num_employees": 256, "country": "UK"},
	}
	got, err := client.Documents.Import(ctx, collectionName, body, nil)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Import_With_Partial_Error(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents/import", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{"success": true}
			{"code":400,"document":"{\"id\": \"2\", \"company_name\": \"Orbit Inc.\", \"num_employees\": 256, \"country\": \"UK\"}","error":"Field `+
			"`company_name`"+` has been declared in the schema, but is not found in the document.","success":false}
		`)
	})

	want := []*ImportDocumentResponse{
		{Success: true},
		{
			Code:     Int(400),
			Document: String("{\"id\": \"2\", \"company_name\": \"Orbit Inc.\", \"num_employees\": 256, \"country\": \"UK\"}"),
			Error:    String("Field `company_name` has been declared in the schema, but is not found in the document."),
			Success:  false,
		},
	}

	ctx := context.Background()
	body := []map[string]interface{}{
		{"id": "1", "company_name": "Stark Industries", "num_employees": 5215, "country": "USA"},
		{"id": "2", "name": "Orbit Inc.", "num_employees": 256, "country": "UK"},
	}
	got, err := client.Documents.Import(ctx, collectionName, body, nil)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestDocumentsService_Search(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"
	u := fmt.Sprintf("/collections/%s/documents/search", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `
			{
				"facet_counts": [],
				"found": 1,
				"hits": [
					{
						"document": {
							"company_name": "Stark Industries",
							"country": "USA",
							"id": "123",
							"num_employees": 5215
						},
						"highlight": {
							"company_name": {
								"matched_tokens": [
									"Stark"
								],
								"snippet": "<mark>Stark</mark> Industries"
							}
						},
						"highlights": [
							{
								"field": "company_name",
								"matched_tokens": [
									"Stark"
								],
								"snippet": "<mark>Stark</mark> Industries"
							}
						],
						"text_match": 578730123365187705,
						"text_match_info": {
							"best_field_score": "1108091338752",
							"best_field_weight": 15,
							"fields_matched": 1,
							"score": "578730123365187705",
							"tokens_matched": 1
						}
					}
				],
				"out_of": 1,
				"page": 1,
				"request_params": {
					"collection_name": "companies",
					"per_page": 10,
					"q": "stark"
				},
				"search_cutoff": false,
				"search_time_ms": 11
			}
		`)
	})

	ctx := context.Background()
	opts := &SearchParameters{
		Q:        "start",
		QueryBy:  "company_name",
		FilterBy: String("num_employees:>100"),
		SortBy:   String("num_employees:desc"),
	}
	got, err := client.Documents.Search(ctx, collectionName, opts)
	assert.NoError(t, err)

	want := &SearchResult{
		FacetCounts: []*FacetCounts{},
		Found:       Int(1),
		Hits: []*SearchResultHit{{
			Document: map[string]interface{}{
				"company_name":  "Stark Industries",
				"country":       "USA",
				"id":            "123",
				"num_employees": float64(5215),
			},
			Highlight: map[string]interface{}{
				"company_name": map[string]interface{}{
					"matched_tokens": []interface{}{"Stark"},
					"snippet":        "<mark>Stark</mark> Industries",
				},
			},
			Highlights: []*SearchHighlight{{
				Field:         String("company_name"),
				MatchedTokens: []string{"Stark"},
				Snippet:       String("<mark>Stark</mark> Industries"),
			}},
			TextMatch: Int64(578730123365187705),
			TextMatchInfo: struct {
				BestFieldScore  string `json:"best_field_score"`
				BestFieldWeight int    `json:"best_field_weight"`
				FieldsMatched   int    `json:"fields_matched"`
				Score           string `json:"score"`
				TokensMatched   int    `json:"tokens_matched"`
			}{
				BestFieldScore:  "1108091338752",
				BestFieldWeight: 15,
				FieldsMatched:   1,
				Score:           "578730123365187705",
				TokensMatched:   1,
			},
		}},
		OutOf: Int(1),
		Page:  Int(1),
		RequestParams: &struct {
			CollectionName string `json:"collection_name"`
			PerPage        int    `json:"per_page"`
			Q              string `json:"q"`
		}{
			CollectionName: "companies",
			PerPage:        10,
			Q:              "stark",
		},
		SearchCutoff: Bool(false),
		SearchTimeMs: Int(11),
	}
	assert.Equal(t, want, got)
}
