package typesense

import (
	"context"
)

type OperationsService service

func (s *OperationsService) Snapshot(ctx context.Context, opts *TakeSnapshotParams) (*SuccessStatus, error) {
	u := "/operations/snapshot"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OperationsService) Vote(ctx context.Context) (*SuccessStatus, error) {
	u := "/operations/vote"
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OperationsService) ClearCache(ctx context.Context) (*SuccessStatus, error) {
	u := "/operations/cache/clear"
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OperationsService) CompactDB(ctx context.Context) (*SuccessStatus, error) {
	u := "/operations/compact/db"
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OperationsService) ResetPeers(ctx context.Context) (*SuccessStatus, error) {
	u := "/operations/reset_peers"
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	res := &SuccessStatus{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
