package typesense

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

const (
	Version = "v0.25.0" // typesense server version

	defaultServerURL = "http://localhost:8108"

	headerAPIKEy      = "X-TYPESENSE-API-KEY"
	headerContentType = "Content-Type"

	defaultMediaType = "application/json"
)

var errNonNilContext = errors.New("context must be non-nil")

type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	ServerURL *url.URL

	common service

	Collections     *CollectionsService
	Documents       *DocumentsService
	Keys            *KeysService
	RateLimits      *RateLimitsService
	Operations      *OperationsService
	Meta            *MetaService
	Overrides       *OverridesService
	Aliases         *AliasesService
	AnalyticsRules  *AnalyticsRulesService
	AnalyticsEvents *AnalyticsEventsService
	Presets         *PresetsService
	Synonyms        *SynonymsService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client, serverURL string) (*Client, error) {
	c := &Client{client: httpClient}

	if c.client == nil {
		c.client = &http.Client{}
	}
	if serverURL != "" {
		serverURL = strings.TrimSuffix(serverURL, "/")	
		url, err := url.Parse(serverURL)
		if err != nil {
			return nil, err
		}
		c.ServerURL = url
	} else {
		c.ServerURL, _ = url.Parse(defaultServerURL)
	}

	c.initialize()
	return c, nil
}

func (c *Client) WithAPIKey(apiKey string) *Client {
	c2 := c.copy()
	defer c2.initialize()
	transport := c2.client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	c2.client.Transport = roundTripperFunc(
		func(r *http.Request) (*http.Response, error) {
			req := r.Clone(r.Context())
			req.Header.Set(headerAPIKEy, apiKey)
			return transport.RoundTrip(req)
		},
	)
	return c2
}

func (c *Client) initialize() {
	c.common.client = c
	c.Collections = (*CollectionsService)(&c.common)
	c.Documents = (*DocumentsService)(&c.common)
	c.Keys = (*KeysService)(&c.common)
	c.RateLimits = (*RateLimitsService)(&c.common)
	c.Operations = (*OperationsService)(&c.common)
	c.Meta = (*MetaService)(&c.common)
	c.Aliases = (*AliasesService)(&c.common)
	c.Overrides = (*OverridesService)(&c.common)
	c.AnalyticsRules = (*AnalyticsRulesService)(&c.common)
	c.AnalyticsEvents = (*AnalyticsEventsService)(&c.common)
	c.Presets  = (*PresetsService)(&c.common)
	c.Synonyms = (*SynonymsService)(&c.common)
}

func (c *Client) copy() *Client {
	c.clientMu.Lock()

	clientCopy := *c.client
	clone := Client{
		client:    &clientCopy,
		ServerURL: c.ServerURL,
	}

	c.clientMu.Unlock()
	return &clone
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

type RequestOption func(req *http.Request)

func (c *Client) NewRequest(method, urlStr string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	u, err := c.ServerURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		switch body := body.(type) {
		case []map[string]interface{}:
			for _, item := range body {
				err := json.NewEncoder(buf).Encode(item)
				if err != nil {
					return nil, err
				}
			}
		case *[]map[string]interface{}:
			for _, item := range *body {
				err := json.NewEncoder(buf).Encode(item)
				if err != nil {
					return nil, err
				}
			}
		default:
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	case *[]map[string]interface{}:
		decoder := json.NewDecoder(resp.Body)
		for {
			var line map[string]interface{}
			if err := decoder.Decode(&line); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			*v = append(*v, line)
		}
	case *[]*ImportDocumentResponse:
		decoder := json.NewDecoder(resp.Body)
		for decoder.More() {
			var line ImportDocumentResponse
			if err := decoder.Decode(&line); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			*v = append(*v, &line)
		}
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}
	return err
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	err = extractApiError(resp)
	if err != nil {
		defer resp.Body.Close()
	}

	return resp, err
}

func extractApiError(r *http.Response) error {
	if c := r.StatusCode; c == 200 || c == 201 {
		return nil
	}

	apiResponse := &ApiResponse{}
	apiError := &ApiError{Response: r}
	body, err := io.ReadAll(r.Body)

	if err == nil && body != nil {
		json.Unmarshal(body, apiResponse)
	}
	apiError.StatusCode = r.StatusCode
	apiError.Body = *apiResponse

	return apiError
}

type ApiError struct {
	Response   *http.Response `json:"-"`
	StatusCode int            `json:"status_code"`
	Body       ApiResponse    `json:"body"`
}

func (r *ApiError) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.StatusCode, r.Body)
}

func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func (c *Client) Ping() error {
	_, err := c.NewRequest("GET", "/health", nil)
	if err != nil {
		return err
	}
	return nil
}
