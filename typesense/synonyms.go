package typesense

import (
	"context"
	"fmt"
)

type SynonymsService service

func (s *SynonymsService) List(ctx context.Context, collectionName string) (*SearchSynonymsResponse, error) {
	u := fmt.Sprintf("/collections/%s/synonyms", collectionName)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SearchSynonymsResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SynonymsService) Get(ctx context.Context, collectionName, synonymId string) (*SearchSynonym, error) {
	u := fmt.Sprintf("/collections/%s/synonyms/%s", collectionName, synonymId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SearchSynonym{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SynonymsService) Upsert(ctx context.Context, collectionName, synonymId string, body *SearchSynonymSchema) (*SearchSynonym, error) {
	u := fmt.Sprintf("/collections/%s/synonyms/%s", collectionName, synonymId)
	req, err := s.client.NewRequest("PUT", u, body)
	if err != nil {
		return nil, err
	}

	res := &SearchSynonym{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeleteSynonymResponse struct {
	ID string `json:"id"`
}

func (s *SynonymsService) Delete(ctx context.Context, collectionName, synonymId string) (*DeleteSynonymResponse, error) {
	u := fmt.Sprintf("/collections/%s/synonyms/%s", collectionName, synonymId)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteSynonymResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
