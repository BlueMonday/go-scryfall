package scryfall

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/ratelimit"
)

const (
	Version = "0.9.0"

	defaultBaseURL      = "https://api.scryfall.com"
	defaultUserAgent    = "go-scryfall/" + Version
	defaultTimeout      = 30 * time.Second
	defaultReqPerSecond = 10

	dateFormat      = "2006-01-02"
	timestampFormat = "2006-01-02T15:04:05.999Z07:00"
)

// ErrMultipleSecrets is returned if both the grant and client secret are set
// when creating a new Scryfall client.
var ErrMultipleSecrets = errors.New("multiple secrets configured")

// Color represents a color in Magic: The Gathering.
type Color string

const (
	// ColorWhite is the white mana color.
	ColorWhite Color = "W"

	// ColorBlue is the blue mana color.
	ColorBlue Color = "U"

	// ColorBlack is the black mana color.
	ColorBlack Color = "B"

	// ColorRed is the red mana color.
	ColorRed Color = "R"

	// ColorGreen is the green mana color.
	ColorGreen Color = "G"
)

// Date is a date returned by the Scryfall API.
type Date struct {
	time.Time
}

// UnmarshalJSON parses a JSON encoded Scryfall date and stores the result.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	// This assumes that all Scryfall dates use the same the timezone as
	// Wizards of the Coast's offices in Renton, Washington.

	parsedTime, err := time.ParseInLocation(dateFormat, s, time.FixedZone("UTC-8", -8*60*60))
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Format(dateFormat))), nil
}

// Timestamp is a timestamp returned by the Scryfall API.
type Timestamp struct {
	time.Time
}

// UnmarshalJSON parses a JSON encoded Scryfall timestamp and stores the result.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	parsedTime, err := time.Parse(timestampFormat, s)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// Error is a Scryfall API error response.
type Error struct {
	Status   int      `json:"status"`
	Code     string   `json:"code"`
	Details  string   `json:"details"`
	Type     *string  `json:"type"`
	Warnings []string `json:"warnings"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Details)
}

type clientOptions struct {
	baseURL      string
	userAgent    string
	clientSecret string
	grantSecret  string
	client       *http.Client
	limiter      ratelimit.Limiter
}

// ClientOption configures the Scryfall API client.
type ClientOption func(*clientOptions)

// WithBaseURL returns an option which overrides the base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(o *clientOptions) {
		o.baseURL = baseURL
	}
}

// WithUserAgent returns an option which overrides the default HTTP user agent.
func WithUserAgent(userAgent string) ClientOption {
	return func(o *clientOptions) {
		o.userAgent = userAgent
	}
}

// WithClientSecret returns an option which sets the client secret. The client
// secret will configure the client to perform requests as the application
// associated with the client secret.
func WithClientSecret(clientSecret string) ClientOption {
	return func(o *clientOptions) {
		o.clientSecret = clientSecret
	}
}

// WithGrantSecret returns an option which sets the grant secret. The grant
// secret will configure the client to perform requests with the rights of the
// grant account.
func WithGrantSecret(grantSecret string) ClientOption {
	return func(o *clientOptions) {
		o.grantSecret = grantSecret
	}
}

// WithHTTPClient returns an option which overrides the default HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.client = client
	}
}

// WithLimiter returns an option which overrides the default rate limiter. A
// nil ratelimiter will disable rate limiting.
func WithLimiter(limiter ratelimit.Limiter) ClientOption {
	return func(o *clientOptions) {
		o.limiter = limiter
	}
}

// Client is a Scryfall API client.
type Client struct {
	baseURL       *url.URL
	userAgent     string
	authorization string

	client  *http.Client
	limiter ratelimit.Limiter
}

// NewClient returns a new Scryfall API client.
func NewClient(options ...ClientOption) (*Client, error) {
	co := &clientOptions{
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		limiter: ratelimit.New(defaultReqPerSecond),
	}
	for _, option := range options {
		option(co)
	}

	if len(co.clientSecret) != 0 && len(co.grantSecret) != 0 {
		return nil, ErrMultipleSecrets
	}

	var authorization string
	if len(co.clientSecret) != 0 {
		authorization = "Bearer " + co.clientSecret
	}
	if len(co.grantSecret) != 0 {
		authorization = "Bearer " + co.grantSecret
	}

	baseURL, err := url.Parse(co.baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:       baseURL,
		userAgent:     co.userAgent,
		authorization: authorization,
		client:        co.client,
		limiter:       co.limiter,
	}
	return c, nil
}

func (c *Client) doReq(ctx context.Context, req *http.Request, respBody interface{}) error {
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")

	if len(c.authorization) != 0 {
		req.Header.Set("Authorization", c.authorization)
	}
	reqWithContext := req.WithContext(ctx)

	if c.limiter != nil {
		c.limiter.Take()
	}

	resp, err := c.client.Do(reqWithContext)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		scryfallErr := &Error{}
		err = decoder.Decode(scryfallErr)
		if err != nil {
			return err
		}

		return scryfallErr
	}

	return decoder.Decode(respBody)
}

func (c *Client) get(ctx context.Context, relativeURL string, respBody interface{}) error {
	absoluteURL, err := c.baseURL.Parse(relativeURL)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, absoluteURL.String(), nil)
	if err != nil {
		return err
	}

	return c.doReq(ctx, req, respBody)
}

func (c *Client) post(ctx context.Context, relativeURL string, reqBody interface{}, respBody interface{}) error {
	absoluteURL, err := c.baseURL.Parse(relativeURL)
	if err != nil {
		return err
	}

	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(http.MethodPost, absoluteURL.String(), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.doReq(ctx, req, respBody)
}

// listResponse represents a requested sequence of other objects (Cards, Sets,
// etc). List objects may be paginated, and also include information about issues
// raised when generating the list.
type listResponse struct {
	// Data is a list of the requested objects, in a specific order.
	Data json.RawMessage `json:"data"`

	// HasMore is true if this List is paginated and there is a page beyond
	// the current page.
	HasMore bool `json:"has_more"`

	// NextPage contains a full API URI to next page if there is a page
	// beyond the current page.
	NextPage *string `json:"next_page"`

	// TotalCards contains the total number of cards found across all pages
	// if this is a list of Card objects.
	TotalCards *int `json:"total_cards"`

	// Warnings is a list of human-readable warnings issued when generating
	// this list, as strings. Warnings are non-fatal issues that the API
	// discovered with your input. In general, they indicate that the List
	// will not contain the all of the information you requested. You should
	// fix the warnings and re-submit your request.
	Warnings []string `json:"warnings"`
}

func (c *Client) listGet(ctx context.Context, url string, v interface{}) error {
	response := &listResponse{}
	err := c.get(ctx, url, response)
	if err != nil {
		return err
	}

	return json.Unmarshal(response.Data, v)
}
