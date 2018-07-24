package scryfall

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://api.scryfall.com"
	defaultTimeout = 30 * time.Second
	userAgent      = "go-scryfall"

	dateFormat = "2006-01-02"
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
		d.Time = time.Time{}
		return nil
	}

	loc, err := time.LoadLocation("Etc/GMT-8")
	if err != nil {
		return err
	}
	parsedTime, err := time.ParseInLocation(dateFormat, s, loc)
	if err != nil {
		return err
	}
	d.Time = parsedTime
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
	clientSecret string
	grantSecret  string
	client       *http.Client
}

// ClientOption configures the Scryfall API client.
type ClientOption func(*clientOptions)

// WithBaseURL returns an option which overrides the base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(o *clientOptions) {
		o.baseURL = baseURL
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

// Client is a Scryfall API client.
type Client struct {
	baseURL       *url.URL
	authorization string

	client *http.Client
}

// NewClient returns a new Scryfall API client.
func NewClient(options ...ClientOption) (*Client, error) {
	co := &clientOptions{
		baseURL: defaultBaseURL,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
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
		authorization: authorization,
		client:        co.client,
	}
	return c, nil
}

func (c *Client) get(ctx context.Context, relativeURL string, v interface{}) error {
	absoluteURL, err := c.baseURL.Parse(relativeURL)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", absoluteURL.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)
	if len(c.authorization) != 0 {
		req.Header.Set("Authorization", c.authorization)
	}

	reqWithContext := req.WithContext(ctx)
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

	return decoder.Decode(v)
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
