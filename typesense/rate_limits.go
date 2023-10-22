package typesense

import (
	"context"
	"fmt"
)

type RateLimitsService service

type RateLimitAction string

const (
	ALLOW    RateLimitAction = "allow"
	BLOCK    RateLimitAction = "block"
	THROTTLE RateLimitAction = "throttle"
)

type RateLimitRuleSchema struct {
	Action                 RateLimitAction `json:"action"`
	ApplyLimitPerEntity    *bool           `json:"apply_limit_per_entity,omitempty"`
	ApiAddresses           []string        `json:"ip_addresses,omitempty"`
	ApiKeys                []string        `json:"api_keys,omitempty"`
	MaxRequests1m          *int            `json:"max_requests_1m,omitempty"`
	MaxRequests1h          *int            `json:"max_requests_1h,omitempty"`
	AutoBan1mThreshold     *int            `json:"auto_ban_1m_threshold,omitempty"`
	AutoBan1mDurationHours *int            `json:"auto_ban_1m_duration_hours,omitempty"`
	Priority               *int            `json:"priority:omitempty"`
}

type RateLimitRule struct {
	ID                  int             `json:"id"`
	Action              RateLimitAction `json:"action"`
	ApplyLimitPerEntity *bool           `json:"apply_limit_per_entity,omitempty"`
	IpAddresses         []string         `json:"ip_addresses,omitempty"`
	ApiKeys             []string        `json:"api_keys,omitempty"`
	MaxRequests         *struct {
		MinuteThreshold *int `json:"minute_threshold,omitempty"`
		HourThreshold   *int `json:"hour_threshold,omitempty"`
	} `json:"max_requests,omitempty"`
	AutoBan1mThreshold     *int `json:"auto_ban_1m_threshold,omitempty"`
	AutoBan1mDurationHours *int `json:"auto_ban_1m_duration_hours,omitempty"`
	Priority               *int `json:"priority,omitempty"`
}

type RateLimitResponse struct {
	Message string         `json:"message"`
	Rule    *RateLimitRule `json:"rule"`
}

type DeleteRateLimitResponse struct {
	ID int `json:"id"`
}

type RateLimitStatus struct {
	StatusId       int     `json:"status_id"`
	ThrottlingFrom int     `json:"throttling_from"`
	ThrottlingTo   int     `json:"throttling_to"`
	IpAddress      *string `json:"ip_address,omitempty"`
	ApiKey         *string `json:"api_key,omitempty"`
}

type RateLimitExceed struct {
	Id           int    `json:"id"` // rate limit rule id
	Ip           string `json:"ip"`
	ApiKey       string `json:"api_key"`
	RequestCount int    `json:"request_count"`
}

type RateLimitStateType string

func (s *RateLimitsService) List(ctx context.Context) ([]*RateLimitRule, error) {
	u := "/limits"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []*RateLimitRule
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) ListActive(ctx context.Context) ([]*RateLimitStatus, error) {
	u := "/limits/active"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []*RateLimitStatus
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) ListExceeds(ctx context.Context) ([]*RateLimitExceed, error) {
	u := "/limits/exceeds"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []*RateLimitExceed
	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) Create(ctx context.Context, body *RateLimitRuleSchema) (*RateLimitResponse, error) {
	u := "/limits"
	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	res := &RateLimitResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) Get(ctx context.Context, id int) (*RateLimitRule, error) {
	u := fmt.Sprintf("/limits/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res := &RateLimitRule{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) Update(ctx context.Context, id int, body *RateLimitRuleSchema) (*RateLimitResponse, error) {
	u := fmt.Sprintf("/limits/%d", id)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	res := &RateLimitResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) Delete(ctx context.Context, id int) (*DeleteRateLimitResponse, error) {
	u := fmt.Sprintf("/limits/%d", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteRateLimitResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) DeleteActive(ctx context.Context, id int) (*DeleteRateLimitResponse, error) {
	u := fmt.Sprintf("/limits/active/%d", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteRateLimitResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RateLimitsService) DeleteExceeds(ctx context.Context, id int) (*DeleteRateLimitResponse, error) {
	u := fmt.Sprintf("/limits/exceeds/%d", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteRateLimitResponse{}
	err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
