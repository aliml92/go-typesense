package typesense

import "context"

func (s *DocumentsService) MultiSearch(ctx context.Context, body *MultiSearchSearchesParameter, opts *MultiSearchParameters) (*MultiSearchResult, error) {
	u := "/multi_search"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &MultiSearchResult{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
