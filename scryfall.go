package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.scryfall.com"
	defaultTimeout = 30 * time.Second
)

// Color represents a color in Magic: The Gathering.
type Color string

const (
	ColorWhite Color = "W"
	ColorBlue  Color = "U"
	ColorBlack Color = "B"
	ColorRed   Color = "R"
	ColorGreen Color = "G"
)

// Layout categorizes the arrangement of card parts, faces, and other bounded
// regions on cards. The layout can be used to programmatically determine which
// other properties on a card you can expect.
type Layout string

const (
	LayoutNormal           Layout = "normal"
	LayoutSplit            Layout = "split"
	LayoutFlip             Layout = "flip"
	LayoutTransform        Layout = "transform"
	LayoutMeld             Layout = "meld"
	LayoutLeveler          Layout = "leveler"
	LayoutPlanar           Layout = "planar"
	LayoutScheme           Layout = "scheme"
	LayoutVanguard         Layout = "vanguard"
	LayoutToken            Layout = "token"
	LayoutDoubleFacedToken Layout = "double_faced_token"
	LayoutEmblem           Layout = "emblem"
	LayoutAugment          Layout = "augment"
	LayoutHost             Layout = "host"
)

// Frame tracks the major edition of the card frame of used for the re/print in
// question. The frame has gone though several major revisions in Magic’s
// lifetime.
type Frame string

const (
	Frame1993   Frame = "1993"
	Frame1997   Frame = "1997"
	Frame2003   Frame = "2003"
	Frame2015   Frame = "2015"
	FrameFuture Frame = "future"
)

// APIError is a Scryfall API error response.
type APIError struct {
	Status   int      `json:"status"`
	Code     string   `json:"code"`
	Details  string   `json:"details"`
	Type     *string  `json:"type"`
	Warnings []string `json:"warnings"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Details)
}

// ListResponse represents a requested sequence of other objects (Cards, Sets,
// etc). List objects may be paginated, and also include information about issues
// raised when generating the list.
type ListResponse struct {
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

type clientOptions struct {
	client *http.Client
}

// ClientOption configures the Scryfall API client.
type ClientOption func(*clientOptions)

// WithHTTPClient returns an option which overrides the default HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.client = client
	}
}

// Client is a Scryfall API client.
type Client struct {
	client *http.Client
}

// NewClient returns a new Scryfall API client.
func NewClient(options ...ClientOption) *Client {
	co := &clientOptions{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	for _, option := range options {
		option(co)
	}

	return &Client{
		client: co.client,
	}
}

func (c *Client) doReq(ctx context.Context, req *http.Request, v interface{}) error {
	reqWithContext := req.WithContext(ctx)
	resp, err := c.client.Do(reqWithContext)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		apiError := &APIError{}
		err = decoder.Decode(apiError)
		if err != nil {
			return err
		}

		return apiError
	}

	return decoder.Decode(v)
}
