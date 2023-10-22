package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `[{
			"name": "companies",
			"fields": [
			  {
				"name": "company_name",
				"type": "string"
			  },
			  {
				"name": "num_employees",
				"type": "int32"
			  },
			  {
				"name": "country",
				"type": "string",
				"facet": true
			  }
			],
			"default_sorting_field": "num_employees"
		}]`)
	})

	ctx := context.Background()
	got, err := client.Collections.List(ctx)
	assert.NoError(t, err)

	want := []*CollectionSchema{{
		Name: "companies",
		Fields: []*Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: Bool(true),
			},
		},
		DefaultSortingField: String("num_employees"),
	}}
	assert.Equal(t, want, got)
}

func TestCollectionsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "companies",
			"fields": [
			  {
				"name": "company_name",
				"type": "string"
			  },
			  {
				"name": "num_employees",
				"type": "int32"
			  },
			  {
				"name": "country",
				"type": "string",
				"facet": true
			  }
			],
			"num_documents": 0,
			"default_sorting_field": "num_employees"
		}`)
	})

	want := &CollectionSchema{
		Name: "companies",
		Fields: []*Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: Bool(true),
			},
		},
		DefaultSortingField: String("num_employees"),
	}
	ctx := context.Background()
	got, err := client.Collections.Create(ctx, want)
	assert.NoError(t, err)

	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Fields, got.Fields)
	assert.Equal(t, want.DefaultSortingField, got.DefaultSortingField)
	assert.Equal(t, int64(0), *got.NumDocuments)
}

func TestCollectionsService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"

	u := fmt.Sprintf("/collections/%s", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "companies",
			"fields": [
			  {
				"name": "company_name",
				"type": "string"
			  },
			  {
				"name": "num_employees",
				"type": "int32"
			  },
			  {
				"name": "country",
				"type": "string",
				"facet": true
			  }
			],
			"num_documents": 0,
			"default_sorting_field": "num_employees"
		}`)
	})

	want := CollectionSchema{
		Name: collectionName,
		Fields: []*Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: Bool(true),
			},
		},
		DefaultSortingField: String("num_employees"),
	}
	ctx := context.Background()
	got, err := client.Collections.Get(ctx, collectionName)
	assert.NoError(t, err)

	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Fields, got.Fields)
	assert.Equal(t, want.DefaultSortingField, got.DefaultSortingField)
}

func TestCollectionsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"

	u := fmt.Sprintf("/collections/%s", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"fields": [
			  {
				"name": "num_employees",
				"drop": true
			  },
			  {
				"name": "company_category",
				"type": "string"
			  }
			]
		}`)
	})

	want := &CollectionUpdateSchema{
		Fields: []*Field{
			{
				Name: "num_employees",
				Drop: Bool(true),
			},
			{
				Name: "company_category",
				Type: "string",
			},
		},
	}
	ctx := context.Background()
	got, err := client.Collections.Update(ctx, collectionName, want)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestCollectionsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	collectionName := "companies"

	u := fmt.Sprintf("/collections/%s", collectionName)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.NotEmpty(t, r.Header.Get(headerAPIKEy))
		fmt.Fprint(w, `{
			"name": "companies",
			"fields": [
			  {
				"name": "company_name",
				"type": "string"
			  },
			  {
				"name": "num_employees",
				"type": "int32"
			  },
			  {
				"name": "country",
				"type": "string",
				"facet": true
			  }
			],
			"num_documents": 0,
			"default_sorting_field": "num_employees"
		}`)
	})

	want := Collection{
		Name: "companies",
		Fields: []*Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: Bool(true),
			},
		},
		NumDocuments:        Int64(0),
		DefaultSortingField: String("num_employees"),
	}
	ctx := context.Background()
	got, err := client.Collections.Delete(ctx, collectionName)
	assert.NoError(t, err)

	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Fields, got.Fields)
	assert.Equal(t, want.DefaultSortingField, got.DefaultSortingField)
	assert.Equal(t, int64(0), *got.NumDocuments)
}
