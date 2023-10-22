# go-typesense
go-typesense is a Go client library for [typesense](https://typesense.org/). Check the [usage](#usage) section or try the [examples](examples/) to see how to work with typesense.

## Features
* Collections
* Documents
* Aliases
* Analytics
* Multi Search
* Keys
* Meta
* Operations
* Overrides
* Presets
* Rate Limits
* Synonyms

## Installation
```bash 
go get github.com/aliml92/go-typesense
```

## Documentation

## Usage
The `go-typesense` package provides a `Client` for accessing typesense server.
```go
	client, _ := typesense.NewClient(nil, "http://localhost:8108")
	client = client.WithAPIKey("xyz")      
```
### Create a collection
```go
	collectionSchema := &typesense.CollectionSchema{
		Name: "companies",
		Fields: []*typesense.Field{
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
				Facet: typesense.Bool(true),
			},
		},
		DefaultSortingField: typesense.String("num_employees"),
	}

	ctx := context.Background()
	collection, err := client.Collections.Create(ctx, collectionSchema)
```
### Index a document
```go
    type Company struct {
        Name         string `json:"company_name"`
        NumEmployees int    `json:"num_employees"`
        Country      string `json:"country"`
    }

	company := &Company{
		Name:         "Tesla",
		NumEmployees: 127_855,
		Country:      "United States",
	}

	indexedDoc, err := client.Documents.Create(ctx, "companies", company)

```
### Search a collection 
```go
	params := &typesense.SearchParameters{
		Q:        "stark",
		QueryBy:  "company_name, country",
		FilterBy: typesense.String("num_employees:>100"),
		SortBy:   typesense.String("num_employees:desc"),
	}

	result, err := client.Documents.Search(ctx, "companies", params)
```
### Manage access to data
The `/keys` API endpoint in Typesense enables the creation of admin keys for overall system control and scoped API keys, allowing precise control over specific operations such as search, thereby providing a robust mechanism for managing data access. For detailed information, please visit [managing access to data](https://typesense.org/docs/guide/data-access-control.html).

#### Bootstrap API Key
The official documentation for the [Bootstrap API Key](https://typesense.org/docs/guide/data-access-control.html#bootstrap-api-key) recommends creating an admin key using the bootstrap API key (the key passed with the --api-key server configuration during cluster creation). This practice enables the rotation and revocation of these admin keys.

Here is how you can achive this with `go-typesense` library. Let's say `xyz` is 
the bootstrap API key
```go
	
	baseClient, _ := typesense.NewClient(nil, "http://localhost:8108")

	// rootClient is admin client with all persmissions and has no expiration.
	rootClient := baseClient.WithAPIKey("xyz")

	// create an admin key with one-month expiration
	ctx := context.Background()
	expiresAt := time.Now().AddDate(0, 1, 0).Unix()

	keySchema := &typesense.ApiKeySchema{
		Actions:     []string{"*"},
		Collections: []string{"*"},
		Description: typesense.String("An admin key with a one-month expiration."),
		ExpiresAt:   &expiresAt,
		Value:       typesense.String("2bRpK9Mx7YsW"),
	}

	key, err := rootClient.Keys.Create(ctx, keySchema)
	if err != nil {
		log.Fatal(err)
	}

    // revocable admin client for day-to-day operations
	adminClient := baseClient.WithAPIKey(*key.Value)
```
#### Search-only API key
```go
	keySchema := &typesense.ApiKeySchema{
		Actions:     []string{"documents:search"},
		Collections: []string{"*"},
		Description: typesense.String("This API key allows searching for documents across all collections and has an expiration set to one month"),
		ExpiresAt:   &expiresAt,
		Value:       typesense.String("9nl4Kn97qsTpC"),
	}
	key, err := adminClient.Keys.Create(ctx, keySchema)
	if err != nil {
		log.Fatal(err)
	}

	// soClient is a Search-only client
	soClient := baseClient.WithAPIKey(*key.Value)
```
>Note: `Client` instances with an API key must be derived from `baseClient` (a client with no API key set).

### Rate limiting
The `/limits`  API endpoint allows setting rate limits based on client's API key
and ip address. Let's apply rate limiting on `soClient` we created above:
```go
	ruleSchema := &typesense.RateLimitRuleSchema{
		Action: typesense.THROTTLE,     // "throttle"
		ApiKeys: []string{"9nl4Kn97qsTpC"},
		MaxRequests1m: typesense.Int(100),
		MaxRequests1h: typesense.Int(-1),
		AutoBan1mThreshold: typesense.Int(1),
		AutoBan1mDurationHours: typesense.Int(1),
	}
	
	res, err := adminClient.RateLimits.Create(ctx, ruleSchema)
```
Now `soClient` can make only 100 search requests per minute.

## Development
### Code structure
The Typesense client library draws inspiration from the structure of [go-github](https://github.com/google/go-github).

The core component is the `Client`, which serves as the foundation for various services, such as `CollectionsService`, `DocumentsService`, `KeysService`, and others. Each of these services encapsulates specific functionality and endpoints, closely aligning with the individual capabilities and use cases of the Typesense API.
