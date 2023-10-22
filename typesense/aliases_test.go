package typesense

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliasesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/aliases", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `[
		{
		  "collection_name": "companies",
		  "name": "companies_alias"
		}
	  ]`)
	})

	want := []*CollectionAlias{{
		CollectionName: "companies",
		Name:           "companies_alias",
	}}

	got, err := client.Aliases.List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestAliasesService_Upsert(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	aliasName := "companies"

	mux.HandleFunc("/aliases/"+aliasName, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		// Assert request body
		fmt.Fprint(w, `{
		"collection_name": "companies",
		"name": "companies_alias" 
	  }`)
	})

	want := &CollectionAlias{
		CollectionName: "companies",
		Name:           "companies_alias",
	}

	body := &CollectionAliasSchema{
		CollectionName: "companies",
	}

	got, err := client.Aliases.Upsert(context.Background(), aliasName, body)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestAliasesService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	aliasName := "companies"

	mux.HandleFunc("/aliases/"+aliasName, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, `{
		"collection_name": "companies",
		"name": "companies_alias"
	  }`)
	})

	want := &CollectionAlias{
		CollectionName: "companies",
		Name:           "companies_alias",
	}

	got, err := client.Aliases.Get(context.Background(), aliasName)

	assert.NoError(t, err)
	assert.Equal(t, want, got)

	// TODO: Add more test cases
}

func TestAliasesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	aliasName := "companies"

	mux.HandleFunc("/aliases/"+aliasName, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		fmt.Fprint(w, `{
		"collection_name": "companies",
		"name": "companies_alias"
	  }`)
	})

	want := &CollectionAlias{
		CollectionName: "companies",
		Name:           "companies_alias",
	}
	ctx := context.Background()
	got, err := client.Aliases.Delete(ctx, aliasName)

	assert.NoError(t, err)
	assert.Equal(t, want, got)

	// TODO: Add more test cases
}
