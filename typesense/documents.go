package typesense

import (
	"context"
	"fmt"
)

type DirtyValuesOptions string

const (
	CoerceOrDrop   DirtyValuesOptions = "coerce_or_drop"
	CoerceOrReject DirtyValuesOptions = "coerce_or_reject"
	Drop           DirtyValuesOptions = "drop"
	Reject         DirtyValuesOptions = "reject"
)

type DocumentsService service

func (s *DocumentsService) Create(ctx context.Context, collectionName string, body interface{}) (interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DocumentsService) Get(ctx context.Context, collectionName, documentId string) (interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UpdateOptions struct {
	FilterBy    string             `url:"filter_by,omitempty"`
	DirtyValues DirtyValuesOptions `url:"dirty_values,omitmepty"`
}

func (s *DocumentsService) Update(ctx context.Context, collectionName, documentId string, body interface{}) (interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	req, err := s.client.NewRequest("PATCH", u, body)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UpdateByQueryResponse struct {
	NumUpdated int `json:"num_updated"`
}

func (s *DocumentsService) UpdateByQuery(ctx context.Context, collectionName, body interface{}, opts *UpdateOptions) (*UpdateByQueryResponse, error) {
	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("PATCH", u, body)
	if err != nil {
		return nil, err
	}

	res := &UpdateByQueryResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DocumentsService) Delete(ctx context.Context, collectionName, documentId string) (interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents/%s", collectionName, documentId)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeleteOptions struct {
	FilterBy  string `url:"filter_by,omitempty"`
	BatchSize int    `url:"batch_size,omitempty"`
}

type DeleteByQueryResponse struct {
	NumDeleted *int `json:"num_deleted"`
}

func (s *DocumentsService) DeleteByQuery(ctx context.Context, collectionName string, opts *DeleteOptions) (*DeleteByQueryResponse, error) {
	u := fmt.Sprintf("/collections/%s/documents", collectionName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteByQueryResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ExportDocumentsParams defines parameters for ExportDocuments.
type ExportDocumentsParams struct {
	// ExcludeFields List of fields from the document to exclude in the search result
	ExcludeFields string `url:"exclude_fields,omitempty"`

	// FilterBy Filter conditions for refining your search results. Separate multiple conditions with &&.
	FilterBy string `url:"filter_by,omitempty"`

	// IncludeFields List of fields from the document to include in the search result
	IncludeFields string `url:"include_fields,omitempty"`
}

// TODO: handle jsonl
func (s *DocumentsService) Export(ctx context.Context, collectionName string, opts *ExportDocumentsParams) ([]map[string]interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents/export", collectionName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []map[string]interface{}
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ImportDocumentsParams struct {
	Action                   *string `json:"action,omitempty"`
	BatchSize                *int    `json:"batch_size,omitempty"`
	DirtyValues              *string `json:"dirty_values,omitempty"`
	RemoteEmbeddingBatchSize *int    `json:"remote_embedding_batch_size,omitempty"`
}

type ImportDocumentResponse struct {
	Code     *int    `json:"code"`
	Success  bool    `json:"success"`
	Document *string `json:"document"`
	Error    *string `json:"error"`
}

// TODO: handle jsonl
func (s *DocumentsService) Import(ctx context.Context, collectionName string, body []map[string]interface{}, opts *ImportDocumentsParams) ([]*ImportDocumentResponse, error) {
	u := fmt.Sprintf("/collections/%s/documents/import", collectionName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	var res []*ImportDocumentResponse
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (s *DocumentsService) ImportFromReader(ctx context.Context, collectionName string, body io.Reader, opts *ImportDocumentsParams) (io.ReadCloser, error) {
// 	u := fmt.Sprintf("/collections/%s/documents/import", collectionName)
// 	u, err := addOptions(u, opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, err := s.client.NewRequest("POST", u, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var res []*ImportDocumentResponse
// 	err = s.client.Do(ctx, req, &res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (s *DocumentsService) Search(ctx context.Context, collectionName string, opts *SearchParameters) (map[string]interface{}, error) {
	u := fmt.Sprintf("/collections/%s/documents/search", collectionName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
