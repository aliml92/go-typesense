package integration

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/aliml92/go-typesense/typesense"
	"github.com/stretchr/testify/require"
)

func TestDocuments_CRUD(t *testing.T) {
	var err error
	collectionName := "laptops"
	ctx := context.Background()

	// Create collection
	createSchema := &typesense.CollectionSchema{
		Name: collectionName,
		Fields: []*typesense.Field{
			{
				Name: "brand",
				Type: "string",
			},
			{
				Name: "model",
				Type: "string",
			},
			{
				Name: "processor",
				Type: "string",
			},
			{
				Name:  "price",
				Type:  "float",
				Facet: typesense.Bool(true),
			},
			{
				Name:  "rating",
				Type:  "int32",
				Facet: typesense.Bool(true),
			},
		},
		DefaultSortingField: typesense.String("price"),
	}

	_, err = client.Collections.Create(ctx, createSchema)
	require.NoError(t, err)

	// simplified laptop struct for tests
	type Laptop struct {
		Brand     string  `json:"brand"`
		Model     string  `json:"model"`
		Processor string  `json:"processor"`
		Price     float32 `json:"price"`
		Rating    int     `json:"rating"`
	}

	product := &Laptop{
		Brand:     "Dell",
		Model:     "Inspiron 15",
		Processor: "Intel Core i5",
		Price:     599.99,
		Rating:    4,
	}

	res, err := client.Documents.Create(ctx, collectionName, product)
	require.NoError(t, err)

	resMap, ok := res.(map[string]interface{})
	require.True(t, ok)

	require.Equal(t, product.Brand, resMap["brand"])
	require.Equal(t, product.Model, resMap["model"])
	require.Equal(t, product.Processor, resMap["processor"])
	require.Equal(t, product.Price, float32(resMap["price"].(float64)))
	require.Equal(t, product.Rating, int(resMap["rating"].(float64)))

	// update document
	documentID, _ := resMap["id"].(string)
	update := struct {
		Price  float32 `json:"price"`
		Rating int     `json:"rating"`
	}{
		Price:  459.99,
		Rating: 3,
	}

	_, err = client.Documents.Update(ctx, collectionName, documentID, update)
	require.NoError(t, err)

	// delete document
	_, err = client.Documents.Delete(ctx, collectionName, documentID)
	require.NoError(t, err)
}

func TestDocuments_Import(t *testing.T) {
	var err error
	collectionName := "products"
	ctx := context.Background()

	// Create collection
	createSchema := &typesense.CollectionSchema{
		Name: collectionName,
		Fields: []*typesense.Field{
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name:  "price",
				Type:  "float",
				Facet: typesense.Bool(true),
			},
			{
				Name:  "rating",
				Type:  "int32",
				Facet: typesense.Bool(true),
			},
			{
				Name: "category",
				Type: "string",
			},
		},
		DefaultSortingField: typesense.String("price"),
	}

	_, err = client.Collections.Create(ctx, createSchema)
	require.NoError(t, err)
	raw, err := os.ReadFile("testdata/ecommerce.jsonl")
	if err != nil {
		require.FailNow(t, err.Error())
	}

	lines := strings.Split(string(raw), "\n")
	var documents []map[string]interface{}

	for _, line := range lines {
		var data map[string]interface{}
		if err = json.Unmarshal([]byte(line), &data); err != nil {
			require.FailNow(t, err.Error())
		}

		documents = append(documents, data)
	}

	res, err := client.Documents.Import(ctx, collectionName, documents, nil)
	require.NoError(t, err)
	for _, r := range res {
		require.True(t, r.Success)
	}
}
