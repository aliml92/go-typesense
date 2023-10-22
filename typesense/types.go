package typesense

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	Api_key_headerScopes = "api_key_header.Scopes"
)

// Defines values for SearchOverrideRuleMatch.
const (
	Contains SearchOverrideRuleMatch = "contains"
	Exact    SearchOverrideRuleMatch = "exact"
)

// Defines values for IndexDocumentParamsAction.
const (
	Upsert IndexDocumentParamsAction = "upsert"
)

// Defines values for ImportDocumentsParamsImportDocumentsParametersDirtyValues.
// const (
// 	CoerceOrDrop   ImportDocumentsParamsImportDocumentsParametersDirtyValues =
// "coerce_or_drop" 	CoerceOrReject
// ImportDocumentsParamsImportDocumentsParametersDirtyValues =
// "coerce_or_reject" 	Drop
// ImportDocumentsParamsImportDocumentsParametersDirtyValues = "drop" 	Reject
//      ImportDocumentsParamsImportDocumentsParametersDirtyValues = "reject"
// )

// ApiKey defines model for ApiKey.
type ApiKey struct {
	Actions     []string `json:"actions"`
	Collections []string `json:"collections"`
	Description string   `json:"description"`
	ExpiresAt   *int64   `json:"expires_at,omitempty"`
	Id          *int64   `json:"id,omitempty"`
	Value       *string  `json:"value,omitempty"`
	ValuePrefix *string  `json:"value_prefix,omitempty"`
}

// ApiKeySchema defines model for ApiKeySchema.
type ApiKeySchema struct {
	Actions     []string `json:"actions"`
	Collections []string `json:"collections"`
	Description *string  `json:"description"`
	ExpiresAt   *int64   `json:"expires_at,omitempty"`
	Value       *string  `json:"value,omitempty"`
}

// ApiKeysResponse defines model for ApiKeysResponse.
type ApiKeysResponse struct {
	Keys []*ApiKey `json:"keys"`
}

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Message string `json:"message"`
}

// Collection defines model for Collection.
type Collection struct {
	// CreatedAt Timestamp of when the collection was created (Unix epoch in
	// seconds)
	CreatedAt *int64 `json:"created_at,omitempty"`

	// DefaultSortingField The name of an int32 / float field that determines
	// the order in which the search results are ranked when a sort_by clause is
	// not provided during searching. This field must indicate some kind of
	// popularity.
	DefaultSortingField *string `json:"default_sorting_field,omitempty"`

	// EnableNestedFields Enables experimental support at a collection level for
	// nested object or object array fields. This field is only available if the
	// Typesense server is version `0.24.0.rcn34` or later.
	EnableNestedFields *bool `json:"enable_nested_fields,omitempty"`

	// Fields A list of fields for querying, filtering and faceting
	Fields []*Field `json:"fields"`

	// Name Name of the collection
	Name string `json:"name"`

	// NumDocuments Number of documents in the collection
	NumDocuments *int64 `json:"num_documents,omitempty"`

	// SymbolsToIndex List of symbols or special characters to be indexed.
	SymbolsToIndex []string `json:"symbols_to_index,omitempty"`

	// TokenSeparators List of symbols or special characters to be used for
	// splitting the text into individual words in addition to space and
	// new-line characters.
	TokenSeparators []string `json:"token_separators,omitempty"`
}

// CollectionSchema defines model for CollectionSchema.
type CollectionSchema struct {
	// DefaultSortingField The name of an int32 / float field that determines
	// the order in which the search results are ranked when a sort_by clause is
	// not provided during searching. This field must indicate some kind of
	// popularity.
	DefaultSortingField *string `json:"default_sorting_field,omitempty"`

	// EnableNestedFields Enables experimental support at a collection level for
	// nested object or object array fields. This field is only available if the
	// Typesense server is version `0.24.0.rcn34` or later.
	EnableNestedFields *bool `json:"enable_nested_fields,omitempty"`

	// Fields A list of fields for querying, filtering and faceting
	Fields []*Field `json:"fields"`

	// Name Name of the collection
	Name string `json:"name"`

	// SymbolsToIndex List of symbols or special characters to be indexed.
	SymbolsToIndex []string `json:"symbols_to_index,omitempty"`

	// TokenSeparators List of symbols or special characters to be used for
	// splitting the text into individual words in addition to space and
	// new-line characters.
	TokenSeparators []string `json:"token_separators,omitempty"`
}

// Field defines model for Field.
type Field struct {
	Drop  *bool `json:"drop,omitempty"`
	Embed *struct {
		From        []string `json:"from"`
		ModelConfig *struct {
			AccessToken  *string `json:"access_token,omitempty"`
			ApiKey       *string `json:"api_key,omitempty"`
			ClientId     *string `json:"client_id,omitempty"`
			ClientSecret *string `json:"client_secret,omitempty"`
			ModelName    *string `json:"model_name"`
			ProjectId    *string `json:"project_id,omitempty"`
		} `json:"model_config"`
	} `json:"embed,omitempty"`
	Facet    *bool   `json:"facet,omitempty"`
	Index    *bool   `json:"index,omitempty"`
	Infix    *bool   `json:"infix,omitempty"`
	Locale   *string `json:"locale,omitempty"`
	Name     string  `json:"name"`
	NumDim   *int    `json:"num_dim,omitempty"`
	Optional *bool   `json:"optional,omitempty"`
	Sort     *bool   `json:"sort,omitempty"`
	Type     string  `json:"type"`
}

// CollectionAlias defines model for CollectionAlias.
type CollectionAlias struct {
	// CollectionName Name of the collection the alias mapped to
	CollectionName string `json:"collection_name"`

	// Name Name of the collection alias
	Name string `json:"name"`
}

// CollectionAliasSchema defines model for CollectionAliasSchema.
type CollectionAliasSchema struct {
	// CollectionName Name of the collection you wish to map the alias to
	CollectionName string `json:"collection_name"`
}

// CollectionAliasesResponse defines model for CollectionAliasesResponse.
type CollectionAliasesResponse struct {
	Aliases []*CollectionAlias `json:"aliases"`
}

// CollectionUpdateSchema defines model for CollectionUpdateSchema.
type CollectionUpdateSchema struct {
	// Fields A list of fields for querying, filtering and faceting
	Fields []*Field `json:"fields"`
}

// FacetCounts defines model for FacetCounts.
type FacetCounts struct {
	Counts *[]struct {
		Count       *int    `json:"count,omitempty"`
		Highlighted *string `json:"highlighted,omitempty"`
		Value       *string `json:"value,omitempty"`
	} `json:"counts,omitempty"`
	FieldName *string `json:"field_name,omitempty"`
	Stats     *struct {
		Avg         *float64 `json:"avg,omitempty"`
		Max         *float64 `json:"max,omitempty"`
		Min         *float64 `json:"min,omitempty"`
		Sum         *float64 `json:"sum,omitempty"`
		TotalValues *int     `json:"total_values,omitempty"`
	} `json:"stats,omitempty"`
}

// MultiSearchCollectionParameters defines model for
// MultiSearchCollectionParameters.
type MultiSearchCollectionParameters struct {
	// CacheTtl The duration (in seconds) that determines how long the search
	// query is cached.  This value can be set on a per-query basis. Default:
	// 60.
	CacheTtl *int `json:"cache_ttl,omitempty"`

	// Collection The collection to search in.
	Collection string `json:"collection"`

	// DropTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to drop the tokens in
	// the query until enough results are found. Tokens that have the least
	// individual hits are dropped first. Set to 0 to disable. Default: 10
	DropTokensThreshold *int `json:"drop_tokens_threshold,omitempty"`

	// EnableOverrides If you have some overrides defined but want to disable
	// all of them during query time, you can do that by setting this parameter
	// to false
	EnableOverrides *bool `json:"enable_overrides,omitempty"`

	// ExcludeFields List of fields from the document to exclude in the search
	// result
	ExcludeFields *string `json:"exclude_fields,omitempty"`

	// ExhaustiveSearch Setting this to true will make Typesense consider all
	// prefixes and typo  corrections of the words in the query without stopping
	// early when enough results are found  (drop_tokens_threshold and
	// typo_tokens_threshold configurations are ignored).
	ExhaustiveSearch *bool `json:"exhaustive_search,omitempty"`

	// FacetBy A list of fields that will be used for faceting your results on.
	// Separate multiple fields with a comma.
	FacetBy *string `json:"facet_by,omitempty"`

	// FacetQuery Facet values that are returned can now be filtered via this
	// parameter. The matching facet text is also highlighted. For example, when
	// faceting by `category`, you can set `facet_query=category:shoe` to return
	// only facet values that contain the prefix "shoe".
	FacetQuery *string `json:"facet_query,omitempty"`

	// FilterBy Filter conditions for refining youropen api validator search
	// results. Separate multiple conditions with &&.
	FilterBy *string `json:"filter_by,omitempty"`

	// GroupBy You can aggregate search results into groups or buckets by
	// specify one or more `group_by` fields. Separate multiple fields with a
	// comma. To group on a particular field, it must be a faceted field.
	GroupBy *string `json:"group_by,omitempty"`

	// GroupLimit Maximum number of hits to be returned for every group. If the
	// `group_limit` is set as `K` then only the top K hits in each group are
	// returned in the response. Default: 3
	GroupLimit *int `json:"group_limit,omitempty"`

	// HiddenHits A list of records to unconditionally hide from search results.
	// A list of `record_id`s to hide. Eg: to hide records with IDs 123 and 456,
	// you'd specify `123,456`. You could also use the Overrides feature to
	// override search results based on rules. Overrides are applied first,
	// followed by `pinned_hits` and finally `hidden_hits`.
	HiddenHits *string `json:"hidden_hits,omitempty"`

	// HighlightAffixNumTokens The number of tokens that should surround the
	// highlighted text on each side. Default: 4
	HighlightAffixNumTokens *int `json:"highlight_affix_num_tokens,omitempty"`

	// HighlightEndTag The end tag used for the highlighted snippets. Default:
	// `</mark>`
	HighlightEndTag *string `json:"highlight_end_tag,omitempty"`

	// HighlightFields A list of custom fields that must be highlighted even if
	// you don't query  for them
	HighlightFields *string `json:"highlight_fields,omitempty"`

	// HighlightFullFields List of fields which should be highlighted fully
	// without snippeting
	HighlightFullFields *string `json:"highlight_full_fields,omitempty"`

	// HighlightStartTag The start tag used for the highlighted snippets.
	// Default: `<mark>`
	HighlightStartTag *string `json:"highlight_start_tag,omitempty"`

	// IncludeFields List of fields from the document to include in the search
	// result
	IncludeFields *string `json:"include_fields,omitempty"`

	// Infix If infix index is enabled for this field, infix searching can be
	// done on a per-field basis by sending a comma separated string parameter
	// called infix to the search query. This parameter can have 3 values; `off`
	// infix search is disabled, which is default `always` infix search is
	// performed along with regular search `fallback` infix search is performed
	// if regular search does not produce results
	Infix *string `json:"infix,omitempty"`

	// MaxExtraPrefix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraPrefix *int `json:"max_extra_prefix,omitempty"`

	// MaxExtraSuffix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraSuffix *int `json:"max_extra_suffix,omitempty"`

	// MaxFacetValues Maximum number of facet values to be returned.
	MaxFacetValues *int `json:"max_facet_values,omitempty"`

	// MinLen1typo Minimum word length for 1-typo correction to be applied.  The
	// value of num_typos is still treated as the maximum allowed typos.
	MinLen1typo *int `json:"min_len_1typo,omitempty"`

	// MinLen2typo Minimum word length for 2-typo correction to be applied.  The
	// value of num_typos is still treated as the maximum allowed typos.
	MinLen2typo *int `json:"min_len_2typo,omitempty"`

	// NumTypos The number of typographical errors (1 or 2) that would be
	// tolerated. Default: 2
	NumTypos *string `json:"num_typos,omitempty"`

	// Page Results from this specific page number would be fetched.
	Page *int `json:"page,omitempty"`

	// PerPage Number of results to fetch per page. Default: 10
	PerPage *int `json:"per_page,omitempty"`

	// PinnedHits A list of records to unconditionally include in the search
	// results at specific positions. An example use case would be to feature or
	// promote certain items on the top of search results. A list of
	// `record_id:hit_position`. Eg: to include a record with ID 123 at Position
	// 1 and another record with ID 456 at Position 5, you'd specify
	// `123:1,456:5`. You could also use the Overrides feature to override
	// search results based on rules. Overrides are applied first, followed by
	// `pinned_hits` and  finally `hidden_hits`.
	PinnedHits *string `json:"pinned_hits,omitempty"`

	// PreSegmentedQuery You can index content from any logographic language
	// into Typesense if you are able to segment / split the text into
	// space-separated words yourself  before indexing and querying.
	// Set this parameter to true to do the same
	PreSegmentedQuery *bool `json:"pre_segmented_query,omitempty"`

	// Prefix Boolean field to indicate that the last word in the query should
	// be treated as a prefix, and not as a whole word. This is used for
	// building autocomplete and instant search interfaces. Defaults to true.
	Prefix *string `json:"prefix,omitempty"`

	// Preset Search using a bunch of search parameters by setting this
	// parameter to the name of the existing Preset.
	Preset *string `json:"preset,omitempty"`

	// PrioritizeExactMatch Set this parameter to true to ensure that an exact
	// match is ranked above the others
	PrioritizeExactMatch *bool `json:"prioritize_exact_match,omitempty"`

	// Q The query text to search for in the collection. Use * as the search
	// string to return all documents. This is typically useful when used in
	// conjunction with filter_by.
	Q *string `json:"q,omitempty"`

	// QueryBy A list of `string` fields that should be queried against.
	// Multiple fields are separated with a comma.
	QueryBy *string `json:"query_by,omitempty"`

	// QueryByWeights The relative weight to give each `query_by` field when
	// ranking results. This can be used to boost fields in priority, when
	// looking for matches. Multiple fields are separated with a comma.
	QueryByWeights *string `json:"query_by_weights,omitempty"`

	// RemoteEmbeddingNumTries Number of times to retry fetching remote
	// embeddings.
	RemoteEmbeddingNumTries *int `json:"remote_embedding_num_tries,omitempty"`

	// RemoteEmbeddingTimeoutMs Timeout (in milliseconds) for fetching remote
	// embeddings.
	RemoteEmbeddingTimeoutMs *int `json:"remote_embedding_timeout_ms,omitempty"`

	// SearchCutoffMs Typesense will attempt to return results early if the
	// cutoff time has elapsed.  This is not a strict guarantee and facet
	// computation is not bound by this parameter.
	SearchCutoffMs *int `json:"search_cutoff_ms,omitempty"`

	// SnippetThreshold Field values under this length will be fully
	// highlighted, instead of showing a snippet of relevant portion. Default:
	// 30
	SnippetThreshold *int `json:"snippet_threshold,omitempty"`

	// SortBy A list of numerical fields and their corresponding sort orders
	// that will be used for ordering your results. Up to 3 sort fields can be
	// specified. The text similarity score is exposed as a special
	// `_text_match` field that you can use in the list of sorting fields. If no
	// `sort_by` parameter is specified, results are sorted by
	// `_text_match:desc,default_sorting_field:desc`
	SortBy *string `json:"sort_by,omitempty"`

	// TypoTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to look for tokens with
	// more typos until enough results are found. Default: 100
	TypoTokensThreshold *int `json:"typo_tokens_threshold,omitempty"`

	// UseCache Enable server side caching of search query results. By default,
	// caching is disabled.
	UseCache *bool `json:"use_cache,omitempty"`

	// VectorQuery Vector query expression for fetching documents "closest" to a
	// given query/document vector.
	VectorQuery *string `json:"vector_query,omitempty"`
}

// MultiSearchParameters Parameters for the multi search API.
type MultiSearchParameters struct {
	// CacheTtl The duration (in seconds) that determines how long the search
	// query is cached.  This value can be set on a per-query basis. Default:
	// 60.
	CacheTtl *int `json:"cache_ttl,omitempty"`

	// DropTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to drop the tokens in
	// the query until enough results are found. Tokens that have the least
	// individual hits are dropped first. Set to 0 to disable. Default: 10
	DropTokensThreshold *int `json:"drop_tokens_threshold,omitempty"`

	// EnableOverrides If you have some overrides defined but want to disable
	// all of them during query time, you can do that by setting this parameter
	// to false
	EnableOverrides *bool `json:"enable_overrides,omitempty"`

	// ExcludeFields List of fields from the document to exclude in the search
	// result
	ExcludeFields *string `json:"exclude_fields,omitempty"`

	// ExhaustiveSearch Setting this to true will make Typesense consider all
	// prefixes and typo  corrections of the words in the query without stopping
	// early when enough results are found  (drop_tokens_threshold and
	// typo_tokens_threshold configurations are ignored).
	ExhaustiveSearch *bool `json:"exhaustive_search,omitempty"`

	// FacetBy A list of fields that will be used for faceting your results on.
	// Separate multiple fields with a comma.
	FacetBy *string `json:"facet_by,omitempty"`

	// FacetQuery Facet values that are returned can now be filtered via this
	// parameter. The matching facet text is also highlighted. For example, when
	// faceting by `category`, you can set `facet_query=category:shoe` to return
	// only facet values that contain the prefix "shoe".
	FacetQuery *string `json:"facet_query,omitempty"`

	// FilterBy Filter conditions for refining youropen api validator search
	// results. Separate multiple conditions with &&.
	FilterBy *string `json:"filter_by,omitempty"`

	// GroupBy You can aggregate search results into groups or buckets by
	// specify one or more `group_by` fields. Separate multiple fields with a
	// comma. To group on a particular field, it must be a faceted field.
	GroupBy *string `json:"group_by,omitempty"`

	// GroupLimit Maximum number of hits to be returned for every group. If the
	// `group_limit` is set as `K` then only the top K hits in each group are
	// returned in the response. Default: 3
	GroupLimit *int `json:"group_limit,omitempty"`

	// HiddenHits A list of records to unconditionally hide from search results.
	// A list of `record_id`s to hide. Eg: to hide records with IDs 123 and 456,
	// you'd specify `123,456`. You could also use the Overrides feature to
	// override search results based on rules. Overrides are applied first,
	// followed by `pinned_hits` and finally `hidden_hits`.
	HiddenHits *string `json:"hidden_hits,omitempty"`

	// HighlightAffixNumTokens The number of tokens that should surround the
	// highlighted text on each side. Default: 4
	HighlightAffixNumTokens *int `json:"highlight_affix_num_tokens,omitempty"`

	// HighlightEndTag The end tag used for the highlighted snippets. Default:
	// `</mark>`
	HighlightEndTag *string `json:"highlight_end_tag,omitempty"`

	// HighlightFields A list of custom fields that must be highlighted even if
	// you don't query  for them
	HighlightFields *string `json:"highlight_fields,omitempty"`

	// HighlightFullFields List of fields which should be highlighted fully
	// without snippeting
	HighlightFullFields *string `json:"highlight_full_fields,omitempty"`

	// HighlightStartTag The start tag used for the highlighted snippets.
	// Default: `<mark>`
	HighlightStartTag *string `json:"highlight_start_tag,omitempty"`

	// IncludeFields List of fields from the document to include in the search
	// result
	IncludeFields *string `json:"include_fields,omitempty"`

	// Infix If infix index is enabled for this field, infix searching can be
	// done on a per-field basis by sending a comma separated string parameter
	// called infix to the search query. This parameter can have 3 values; `off`
	// infix search is disabled, which is default `always` infix search is
	// performed along with regular search `fallback` infix search is performed
	// if regular search does not produce results
	Infix *string `json:"infix,omitempty"`

	// MaxExtraPrefix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraPrefix *int `json:"max_extra_prefix,omitempty"`

	// MaxExtraSuffix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraSuffix *int `json:"max_extra_suffix,omitempty"`

	// MaxFacetValues Maximum number of facet values to be returned.
	MaxFacetValues *int `json:"max_facet_values,omitempty"`

	// MinLen1typo Minimum word length for 1-typo correction to be applied.  The
	// value of num_typos is still treated as the maximum allowed typos.
	MinLen1typo *int `json:"min_len_1typo,omitempty"`

	// MinLen2typo Minimum word length for 2-typo correction to be applied.  The
	// value of num_typos is still treated as the maximum allowed typos.
	MinLen2typo *int `json:"min_len_2typo,omitempty"`

	// NumTypos The number of typographical errors (1 or 2) that would be
	// tolerated. Default: 2
	NumTypos *string `json:"num_typos,omitempty"`

	// Page Results from this specific page number would be fetched.
	Page *int `json:"page,omitempty"`

	// PerPage Number of results to fetch per page. Default: 10
	PerPage *int `json:"per_page,omitempty"`

	// PinnedHits A list of records to unconditionally include in the search
	// results at specific positions. An example use case would be to feature or
	// promote certain items on the top of search results. A list of
	// `record_id:hit_position`. Eg: to include a record with ID 123 at Position
	// 1 and another record with ID 456 at Position 5, you'd specify
	// `123:1,456:5`. You could also use the Overrides feature to override
	// search results based on rules. Overrides are applied first, followed by
	// `pinned_hits` and  finally `hidden_hits`.
	PinnedHits *string `json:"pinned_hits,omitempty"`

	// PreSegmentedQuery You can index content from any logographic language
	// into Typesense if you are able to segment / split the text into
	// space-separated words yourself  before indexing and querying.
	// Set this parameter to true to do the same
	PreSegmentedQuery *bool `json:"pre_segmented_query,omitempty"`

	// Prefix Boolean field to indicate that the last word in the query should
	// be treated as a prefix, and not as a whole word. This is used for
	// building autocomplete and instant search interfaces. Defaults to true.
	Prefix *string `json:"prefix,omitempty"`

	// Preset Search using a bunch of search parameters by setting this
	// parameter to the name of the existing Preset.
	Preset *string `json:"preset,omitempty"`

	// PrioritizeExactMatch Set this parameter to true to ensure that an exact
	// match is ranked above the others
	PrioritizeExactMatch *bool `json:"prioritize_exact_match,omitempty"`

	// Q The query text to search for in the collection. Use * as the search
	// string to return all documents. This is typically useful when used in
	// conjunction with filter_by.
	Q *string `json:"q,omitempty"`

	// QueryBy A list of `string` fields that should be queried against.
	// Multiple fields are separated with a comma.
	QueryBy *string `json:"query_by,omitempty"`

	// QueryByWeights The relative weight to give each `query_by` field when
	// ranking results. This can be used to boost fields in priority, when
	// looking for matches. Multiple fields are separated with a comma.
	QueryByWeights *string `json:"query_by_weights,omitempty"`

	// RemoteEmbeddingNumTries Number of times to retry fetching remote
	// embeddings.
	RemoteEmbeddingNumTries *int `json:"remote_embedding_num_tries,omitempty"`

	// RemoteEmbeddingTimeoutMs Timeout (in milliseconds) for fetching remote
	// embeddings.
	RemoteEmbeddingTimeoutMs *int `json:"remote_embedding_timeout_ms,omitempty"`

	// SearchCutoffMs Typesense will attempt to return results early if the
	// cutoff time has elapsed.  This is not a strict guarantee and facet
	// computation is not bound by this parameter.
	SearchCutoffMs *int `json:"search_cutoff_ms,omitempty"`

	// SnippetThreshold Field values under this length will be fully
	// highlighted, instead of showing a snippet of relevant portion. Default:
	// 30
	SnippetThreshold *int `json:"snippet_threshold,omitempty"`

	// SortBy A list of numerical fields and their corresponding sort orders
	// that will be used for ordering your results. Up to 3 sort fields can be
	// specified. The text similarity score is exposed as a special
	// `_text_match` field that you can use in the list of sorting fields. If no
	// `sort_by` parameter is specified, results are sorted by
	// `_text_match:desc,default_sorting_field:desc`
	SortBy *string `json:"sort_by,omitempty"`

	// TypoTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to look for tokens with
	// more typos until enough results are found. Default: 100
	TypoTokensThreshold *int `json:"typo_tokens_threshold,omitempty"`

	// UseCache Enable server side caching of search query results. By default,
	// caching is disabled.
	UseCache *bool `json:"use_cache,omitempty"`

	// VectorQuery Vector query expression for fetching documents "closest" to a
	// given query/document vector.
	VectorQuery *string `json:"vector_query,omitempty"`
}

// MultiSearchResult defines model for MultiSearchResult.
// type MultiSearchResult struct {
// 	Results []*SearchResult `json:"results"`
// }

type MultiSearchResult struct {
	Results []ResultOrError `json:"results"`
}

type ResultOrError struct {
	SearchResult *SearchResult
	SearchError  *SearchError
}

type SearchError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (roe *ResultOrError) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, []byte(`code`)) {
		var ae SearchError
		if err := json.Unmarshal(data, &ae); err == nil {
			roe.SearchError = &ae
			return nil
		}
	}
	var sr SearchResult
	if err := json.Unmarshal(data, &sr); err == nil {
		roe.SearchResult = &sr
		return nil
	}

	return fmt.Errorf("data does not match SearchResult or SearchError")
}

// MultiSearchSearchesParameter defines model for MultiSearchSearchesParameter.
type MultiSearchSearchesParameter struct {
	Searches []MultiSearchCollectionParameters `json:"searches"`
}

// SearchGroupedHit defines model for SearchGroupedHit.
type SearchGroupedHit struct {
	GroupKey []string `json:"group_key"`

	// Hits The documents that matched the search query
	Hits []SearchResultHit `json:"hits"`
}

// SearchHighlight defines model for SearchHighlight.
type SearchHighlight struct {
	Field *string `json:"field,omitempty"`

	// Indices The indices property will be present only for string[] fields and
	// will contain the corresponding indices of the snippets in the search
	// field
	Indices       []int    `json:"indices,omitempty"`
	MatchedTokens []string `json:"matched_tokens,omitempty"`

	// Snippet Present only for (non-array) string fields
	Snippet *string `json:"snippet,omitempty"`

	// Snippets Present only for (array) string[] fields
	Snippets []string `json:"snippets,omitempty"`

	// Value Full field value with highlighting, present only for (non-array)
	// string fields
	Value *string `json:"value,omitempty"`

	// Values Full field value with highlighting, present only for (array)
	// string[] fields
	Values []string `json:"values,omitempty"`
}

// SearchOverride defines model for SearchOverride.
type SearchOverride struct {
	// Excludes List of document `id`s that should be excluded from the search
	// results.
	Excludes []*SearchOverrideExclude `json:"excludes,omitempty"`

	// FilterBy A filter by clause that is applied to any search query that
	// matches the override rule.
	FilterBy *string `json:"filter_by,omitempty"`
	Id       *string `json:"id,omitempty"`

	// Includes List of document `id`s that should be included in the search
	// results with their corresponding `position`s.
	Includes []*SearchOverrideInclude `json:"includes,omitempty"`

	// RemoveMatchedTokens Indicates whether search query tokens that exist in
	// the override's rule should be removed from the search query.
	RemoveMatchedTokens *bool              `json:"remove_matched_tokens,omitempty"`
	Rule                SearchOverrideRule `json:"rule"`
}

// SearchOverrideExclude defines model for SearchOverrideExclude.
type SearchOverrideExclude struct {
	// Id document id that should be excluded from the search results.
	Id string `json:"id"`
}

// SearchOverrideInclude defines model for SearchOverrideInclude.
type SearchOverrideInclude struct {
	// Id document id that should be included
	Id string `json:"id"`

	// Position position number where document should be included in the search
	// results
	Position int `json:"position"`
}

// SearchOverrideRule defines model for SearchOverrideRule.
type SearchOverrideRule struct {
	// Match Indicates whether the match on the query term should be `exact` or
	// `contains`. If we want to match all queries that contained the word
	// `apple`, we will use the `contains` match instead.
	Match SearchOverrideRuleMatch `json:"match"`

	// Query Indicates what search queries should be overridden
	Query string `json:"query"`
}

// SearchOverrideRuleMatch Indicates whether the match on the query term should
// be `exact` or `contains`. If we want to match all queries that contained the
// word `apple`, we will use the `contains` match instead.
type SearchOverrideRuleMatch string

// SearchOverrideSchema defines model for SearchOverrideSchema.
type SearchOverrideSchema struct {
	// Excludes List of document `id`s that should be excluded from the search
	// results.
	Excludes *[]SearchOverrideExclude `json:"excludes,omitempty"`

	// FilterBy A filter by clause that is applied to any search query that
	// matches the override rule.
	FilterBy *string `json:"filter_by,omitempty"`

	// Includes List of document `id`s that should be included in the search
	// results with their corresponding `position`s.
	Includes *[]SearchOverrideInclude `json:"includes,omitempty"`

	// RemoveMatchedTokens Indicates whether search query tokens that exist in
	// the override's rule should be removed from the search query.
	RemoveMatchedTokens *bool              `json:"remove_matched_tokens,omitempty"`
	Rule                SearchOverrideRule `json:"rule"`
}

// SearchOverridesResponse defines model for SearchOverridesResponse.
type SearchOverridesResponse struct {
	Overrides []*SearchOverride `json:"overrides"`
}

type SearchParameters struct {
	// 1. Query params

	// Q The query text to search for in the collection. Use * as the search
	// string to return all documents. This is typically useful when used in
	// conjunction with filter_by.
	Q string `url:"q"`

	// QueryBy A list of `string` fields that should be queried against.
	// Multiple fields are separated with a comma.
	QueryBy string `url:"query_by"`

	// Prefix Boolean field to indicate that the last word in the query should
	// be treated as a prefix, and not as a whole word. This is used for
	// building autocomplete and instant search interfaces. Defaults to true.
	Prefix *string `url:"prefix,omitempty"`

	// Infix If infix index is enabled for this field, infix searching can be
	// done on a per-field basis by sending a comma separated string parameter
	// called infix to the search query. This parameter can have 3 values; `off`
	// infix search is disabled, which is default `always` infix search is
	// performed along with regular search `fallback` infix search is performed
	// if regular search does not produce results
	Infix *string `url:"infix,omitempty"`

	// MaxExtraPrefix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraPrefix *int `url:"max_extra_prefix,omitempty"`

	// MaxExtraSuffix There are also 2 parameters that allow you to control the
	// extent of infix searching max_extra_prefix and max_extra_suffix which
	// specify the maximum number of symbols before or after the query that can
	// be present in the token. For example query "K2100" has 2 extra symbols in
	// "6PK2100". By default, any number of prefixes/suffixes can be present for
	// a match.
	MaxExtraSuffix *int `url:"max_extra_suffix,omitempty"`

	// PreSegmentedQuery You can index content from any logographic language
	// into Typesense if you are able to segment / split the text into
	// space-separated words yourself  before indexing and querying.
	// Set this parameter to true to do the same
	PreSegmentedQuery *bool `url:"pre_segmented_query,omitempty"`

	// Preset Search using a bunch of search parameters by setting this
	// parameter to the name of the existing Preset.
	Preset *string `url:"preset,omitempty"`

	// 2. Filter params

	// FilterBy Filter conditions for refining youropen api validator search
	// results. Separate multiple conditions with &&.
	FilterBy *string `url:"filter_by,omitmepty"`

	// 3. Ranking and Sorting params

	// QueryByWeights The relative weight to give each `query_by` field when
	// ranking results. This can be used to boost fields in priority, when
	// looking for matches. Multiple fields are separated with a comma.
	QueryByWeights *string `url:"query_by_weights,omitempty"`

	// In a multi-field matching context, this parameter determines how the 
	// representative text match score of a record is calculated.
	// Possible values: `max_score` (default) or `max_weight`.
	TextMatchType *string `url:"text_match_type,omitempty"`

	// SortBy A list of numerical fields and their corresponding sort orders
	// that will be used for ordering your results. Up to 3 sort fields can be
	// specified. The text similarity score is exposed as a special
	// `_text_match` field that you can use in the list of sorting fields. If no
	// `sort_by` parameter is specified, results are sorted by
	// `_text_match:desc,default_sorting_field:desc`
	SortBy *string `url:"sort_by,omitempty"`

	// PrioritizeExactMatch Set this parameter to true to ensure that an exact
	// match is ranked above the others
	PrioritizeExactMatch *bool `url:"prioritize_exact_match,omitempty"`

	// PrioritizeTokenPosition Make Typesense prioritize documents where the
	// query words appear earlier in the text.
	PrioritizeTokenPosition *bool `url:"prioritize_token_position,omitempty"`

	// PinnedHits A list of records to unconditionally include in the search
	// results at specific positions. An example use case would be to feature or
	// promote certain items on the top of search results. A list of
	// `record_id:hit_position`. Eg: to include a record with ID 123 at Position
	// 1 and another record with ID 456 at Position 5, you'd specify
	// `123:1,456:5`. You could also use the Overrides feature to override
	// search results based on rules. Overrides are applied first, followed by
	// `pinned_hits` and  finally `hidden_hits`.
	PinnedHits *string `url:"pinned_hits,omitempty"`

	// HiddenHits A list of records to unconditionally hide from search results.
	// A list of `record_id`s to hide. Eg: to hide records with IDs 123 and 456,
	// you'd specify `123,456`. You could also use the Overrides feature to
	// override search results based on rules. Overrides are applied first,
	// followed by `pinned_hits` and finally `hidden_hits`.
	HiddenHits *string `url:"hidden_hits,omitempty"`

	// EnableOverrides If you have some overrides defined but want to disable
	// all of them during query time, you can do that by setting this parameter
	// to false
	EnableOverrides *bool `url:"enable_overrides,omitmepty"`

	// 4. Pagination params

	// Page Results from this specific page number would be fetched.
	Page *int `url:"page,omitempty"`

	// PerPage Number of results to fetch per page. Default: 10
	PerPage *int `url:"per_page,omitempty"`
	
	// Identifies the starting point to return hits from a result set. 
	// Can be used as an alternative to the page parameter.
	Offset *int `url:"offset,omitempty"`

	// Number of hits to fetch. Can be used as an alternative to the per_page 
	// parameter. Default: 10
	Limit *int `url:"limit,omitempty"`

	// 5. Faceting params

	// FacetBy A list of fields that will be used for faceting your results on.
	// Separate multiple fields with a comma.
	FacetBy *string `url:"facet_by,omitempty"`

	// MaxFacetValues Maximum number of facet values to be returned.
	MaxFacetValues *int `url:"max_facet_values,omitempty"`

	// FacetQuery Facet values that are returned can now be filtered via this
	// parameter. The matching facet text is also highlighted. For example, when
	// faceting by `category`, you can set `facet_query=category:shoe` to return
	// only facet values that contain the prefix "shoe".
	FacetQuery *string `url:"facet_query,omitempty"`

	FacetQueryNumTypes *int `url:"facet_query_num_typos,omitempty"`

	// 6. Grouping params

	// GroupBy You can aggregate search results into groups or buckets by
	// specify one or more `group_by` fields. Separate multiple fields with a
	// comma. To group on a particular field, it must be a faceted field.
	GroupBy *string `url:"group_by,omitempty"`

	// GroupLimit Maximum number of hits to be returned for every group. If the
	// `group_limit` is set as `K` then only the top K hits in each group are
	// returned in the response. Default: 3
	GroupLimit *int `url:"group_limit,omitempty"`

	// 7. Result params

	// IncludeFields List of fields from the document to include in the search
	// result
	IncludeFields *string `url:"include_fields,omitempty"`

	// ExcludeFields List of fields from the document to exclude in the search
	// result
	ExcludeFields *string `url:"exclude_fields,omitempty"`

	// HighlightFields A list of custom fields that must be highlighted even if
	// you don't query  for them
	HighlightFields *string `url:"highlight_fields,omitempty"`

	// HighlightFullFields List of fields which should be highlighted fully
	// without snippeting
	HighlightFullFields *string `url:"highlight_full_fields,omitempty"`

	// HighlightAffixNumTokens The number of tokens that should surround the
	// highlighted text on each side. Default: 4
	HighlightAffixNumTokens *int `url:"highlight_affix_num_tokens,omitempty"`

	// HighlightStartTag The start tag used for the highlighted snippets.
	// Default: `<mark>`
	HighlightStartTag *string `url:"highlight_start_tag,omitempty"`

	// HighlightEndTag The end tag used for the highlighted snippets. 
	// Default: `</mark>`
	HighlightEndTag *string `url:"highlight_end_tag,omitempty"`

	// EnableHighlightV1 Flag for enabling/disabling the deprecated, 
	// old highlight structure in the response. Default: true
	EnableHighlightV1 *bool `url:"enable_highlight_v1,omitempty"`

	// SnippetThreshold Field values under this length will be fully
	// highlighted, instead of showing a snippet of relevant portion. Default: 30
	SnippetThreshold *int `url:"snippet_threshold,omitempty"`

	// Maximum number of hits that can be fetched from the collection. 
	// `page` * `per_page` should be less than this number for the search request 
	// to return results. Default: no limit
	LimitHits *int `url:"limit_hits,omitempty"`

	// SearchCutoffMs Typesense will attempt to return results early if the
	// cutoff time has elapsed.  This is not a strict guarantee and facet
	// computation is not bound by this parameter.
	SearchCutoffMs *int `url:"search_cutoff_ms,omitempty"`

	// MaxCandidates Control the number of words that Typesense considers for
	// typo and prefix searching.
	MaxCandidates *int `url:"max_candidates,omitempty"`

	// ExhaustiveSearch Setting this to true will make Typesense consider all
	// prefixes and typo  corrections of the words in the query without stopping
	// early when enough results are found  (drop_tokens_threshold and
	// typo_tokens_threshold configurations are ignored).
	ExhaustiveSearch *bool `url:"exhaustive_search,omitempty"`

	// 8. Type-Tolerance params

	// NumTypos The number of typographical errors (1 or 2) that would be tolerated. 
	// Default: 2
	NumTypos *string `url:"num_typos,omitempty"`

	// MinLen1typo Minimum word length for 1-typo correction to be applied.  
	// The value of num_typos is still treated as the maximum allowed typos.
	MinLen1typo *int `url:"min_len_1typo,omitempty"`

	// MinLen2typo Minimum word length for 2-typo correction to be applied.  
	// The value of num_typos is still treated as the maximum allowed typos.
	MinLen2typo *int `url:"min_len_2type,omitempty"`

	// SplitJoinTokens Treat space as typo: search for q=basket ball if
	// q=basketball is not found or vice-versa. Splitting/joining of tokens will
	// only be attempted if the original query produces no results. To always
	// trigger this behavior, set value to `always`. To disable, set value to
	// `off`. Default is `fallback`.
	SplitJoinTokens *string `url:"split_join_tokes,omitempty"`

	// TypoTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to look for tokens with
	// more typos until enough results are found. Default: 100
	TypoTokensThreshold *int `url:"typo_tokens_threshold,omitempty"`

	// DropTokensThreshold If the number of results found for a specific query
	// is less than this number, Typesense will attempt to drop the tokens in
	// the query until enough results are found. Tokens that have the least
	// individual hits are dropped first. Set to 0 to disable. Default: 10
	DropTokensThreshold *int `url:"drop_tokens_threshold,omitempty"`

	// 9. Caching params

	// UseCache Enable server side caching of search query results. 
	// By default, caching is disabled.
	UseCache *bool `url:"use_cache,omitempty"`

	// CacheTtl The duration (in seconds) that determines how long the search
	// query is cached. This value can be set on a per-query basis. Default: 60.
	CacheTtl *int `url:"cache_ttl,omitempty"`

	// 10. Remote embedding params

	// RemoteEmbeddingNumTries Number of times to retry fetching remote
	// embeddings.
	RemoteEmbeddingNumTries *int `url:"remote_embedding_num_tries,omitempty"`

	// RemoteEmbeddingTimeoutMs Timeout (in milliseconds) for fetching remote
	// embeddings.
	RemoteEmbeddingTimeoutMs *int `url:"remote_embedding_timeout_ms,omitempty"`

	// 11. Other params

	// VectorQuery Vector query expression for fetching documents "closest" to a
	// given query/document vector.
	VectorQuery *string `url:"vector_query,omitempty"`
}

// SearchResult defines model for SearchResult.
type SearchResult struct {
	FacetCounts []*FacetCounts `json:"facet_counts,omitempty"`

	// Found The number of documents found
	Found       *int                `json:"found,omitempty"`
	GroupedHits []*SearchGroupedHit `json:"grouped_hits,omitempty"`

	// Hits The documents that matched the search query
	Hits []*SearchResultHit `json:"hits,omitempty"`

	// OutOf The total number of documents in the collection
	OutOf *int `json:"out_of,omitempty"`

	// Page The search result page number
	Page          *int `json:"page,omitempty"`
	RequestParams *struct {
		CollectionName string `json:"collection_name"`
		PerPage        int    `json:"per_page"`
		Q              string `json:"q"`
	} `json:"request_params,omitempty"`

	// SearchCutoff Whether the search was cut off
	SearchCutoff *bool `json:"search_cutoff,omitempty"`

	// SearchTimeMs The number of milliseconds the search took
	SearchTimeMs *int `json:"search_time_ms,omitempty"`
}

// SearchResultHit defines model for SearchResultHit.
type SearchResultHit struct {
	// Document Can be any key-value pair
	Document map[string]interface{} `json:"document,omitempty"`

	// GeoDistanceMeters Can be any key-value pair
	GeoDistanceMeters map[string]int `json:"geo_distance_meters,omitempty"`

	// Highlight Highlighted version of the matching document
	Highlight map[string]interface{} `json:"highlight,omitempty"`

	// Highlights (Deprecated) Contains highlighted portions of the search
	// fields
	Highlights    []*SearchHighlight `json:"highlights,omitempty"`
	TextMatch     *int64             `json:"text_match,omitempty"`
	TextMatchInfo struct {
		BestFieldScore  string `json:"best_field_score"`
		BestFieldWeight int    `json:"best_field_weight"`
		FieldsMatched   int    `json:"fields_matched"`
		Score           string `json:"score"`
		TokensMatched   int    `json:"tokens_matched"`
	} `json:"text_match_info"`

	// VectorDistance Distance between the query vector and matching document's
	// vector value
	VectorDistance *float32 `json:"vector_distance,omitempty"`
}

// SearchSynonym defines model for SearchSynonym.
type SearchSynonym struct {
	Id *string `json:"id,omitempty"`

	// Root For 1-way synonyms, indicates the root word that words in the
	// `synonyms` parameter map to.
	Root *string `json:"root,omitempty"`

	// Synonyms Array of words that should be considered as synonyms.
	Synonyms []string `json:"synonyms"`
}

// SearchSynonymSchema defines model for SearchSynonymSchema.
type SearchSynonymSchema struct {
	// Root For 1-way synonyms, indicates the root word that words in the
	// `synonyms` parameter map to.
	Root *string `json:"root,omitempty"`

	// Synonyms Array of words that should be considered as synonyms.
	Synonyms []string `json:"synonyms"`
}

// SearchSynonymsResponse defines model for SearchSynonymsResponse.
type SearchSynonymsResponse struct {
	Synonyms []*SearchSynonym `json:"synonyms"`
}

// SuccessStatus defines model for SuccessStatus.
type SuccessStatus struct {
	Success bool `json:"success"`
}

// DeleteDocumentsParams defines parameters for DeleteDocuments.
type DeleteDocumentsParams struct {
	DeleteDocumentsParameters *struct {
		// BatchSize Batch size parameter controls the number of documents that
		// should be deleted at a time. A larger value will speed up deletions,
		// but will impact performance of other operations running on the
		// server.
		BatchSize *int    `json:"batch_size,omitempty"`
		FilterBy  *string `json:"filter_by,omitempty"`
	} `form:"deleteDocumentsParameters,omitempty" json:"deleteDocumentsParameters,omitempty"`
}

// UpdateDocumentsJSONBody defines parameters for UpdateDocuments.
type UpdateDocumentsJSONBody = interface{}

// UpdateDocumentsParams defines parameters for UpdateDocuments.
type UpdateDocumentsParams struct {
	UpdateDocumentsParameters *struct {
		FilterBy *string `json:"filter_by,omitempty"`
	} `form:"updateDocumentsParameters,omitempty" json:"updateDocumentsParameters,omitempty"`
}

// IndexDocumentJSONBody defines parameters for IndexDocument.
type IndexDocumentJSONBody = interface{}

// IndexDocumentParams defines parameters for IndexDocument.
type IndexDocumentParams struct {
	// Action Additional action to perform
	Action *IndexDocumentParamsAction `form:"action,omitempty" json:"action,omitempty"`
}

// IndexDocumentParamsAction defines parameters for IndexDocument.
type IndexDocumentParamsAction string

// ImportDocumentsParamsImportDocumentsParametersDirtyValues defines parameters
// for ImportDocuments.
type ImportDocumentsParamsImportDocumentsParametersDirtyValues string

// SearchCollectionParams defines parameters for SearchCollection.
type SearchCollectionParams struct {
	SearchParameters SearchParameters `form:"searchParameters" json:"searchParameters"`
}

// UpdateDocumentJSONBody defines parameters for UpdateDocument.
type UpdateDocumentJSONBody = interface{}

// MultiSearchParams defines parameters for MultiSearch.
type MultiSearchParams struct {
	MultiSearchParameters MultiSearchParameters `form:"multiSearchParameters" json:"multiSearchParameters"`
}

// TakeSnapshotParams defines parameters for TakeSnapshot.
type TakeSnapshotParams struct {
	// SnapshotPath The directory on the server where the snapshot should be
	// saved.
	SnapshotPath string `form:"snapshot_path" json:"snapshot_path"`
}

// UpsertAliasJSONRequestBody defines body for UpsertAlias for application/json
// ContentType.
type UpsertAliasJSONRequestBody = CollectionAliasSchema

// CreateCollectionJSONRequestBody defines body for CreateCollection for
// application/json ContentType.
type CreateCollectionJSONRequestBody = CollectionSchema

// UpdateCollectionJSONRequestBody defines body for UpdateCollection for
// application/json ContentType.
type UpdateCollectionJSONRequestBody = CollectionUpdateSchema

// UpdateDocumentsJSONRequestBody defines body for UpdateDocuments for
// application/json ContentType.
type UpdateDocumentsJSONRequestBody = UpdateDocumentsJSONBody

// IndexDocumentJSONRequestBody defines body for IndexDocument for
// application/json ContentType.
type IndexDocumentJSONRequestBody = IndexDocumentJSONBody

// UpdateDocumentJSONRequestBody defines body for UpdateDocument for
// application/json ContentType.
type UpdateDocumentJSONRequestBody = UpdateDocumentJSONBody

// UpsertSearchOverrideJSONRequestBody defines body for UpsertSearchOverride for
// application/json ContentType.
type UpsertSearchOverrideJSONRequestBody = SearchOverrideSchema

// UpsertSearchSynonymJSONRequestBody defines body for UpsertSearchSynonym for
// application/json ContentType.
type UpsertSearchSynonymJSONRequestBody = SearchSynonymSchema

// CreateKeyJSONRequestBody defines body for CreateKey for application/json
// ContentType.
type CreateKeyJSONRequestBody = ApiKeySchema

// MultiSearchJSONRequestBody defines body for MultiSearch for application/json
// ContentType.
type MultiSearchJSONRequestBody = MultiSearchSearchesParameter
