package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
)

// Set is an object which represents a group of related Magic cards. All Card
// objects on Scryfall belong to exactly one set.
type Set struct {
	// Code is the unique three or four-letter code for this set.
	Code string `json:"code"`

	// MTGOCode is the unique code for this set on MTGO, which may differ
	// from the regular code.
	MTGOCode string `json:"mtgo_code"`

	// Name is the English name of the set.
	Name string `json:"name"`

	// SetType is a computer-readable classification for this set
	SetType string `json:"set_type"`

	// ReleasedAt is the date the set was released (in GMT-8 Pacific
	// time). Not all sets have a known release date.
	ReleasedAt *string `json:"released_at"`

	// BlockCode is the block code for this set, if any.
	BlockCode *string `json:"block_code"`

	// Block the block or group name code for this set, if any.
	Block *string `json:"block"`

	// ParentSetCode is the set code for the parent set, if any. promo and
	// token sets often have a parent set.
	ParentSetCode string `json:"parent_set_code"`

	// CardCount is the number of cards in this set.
	CardCount int `json:"card_count"`

	// Digital is true if this set was only released on Magic Online.
	Digital bool `json:"digital"`

	// Foil is true if this set contains only foil cards.
	Foil bool `json:"foil"`

	// IconSVGURL is a URI to an SVG file for this set’s icon on Scryfall’s
	// CDN. Hotlinking this image isn’t recommended, because it may change
	// slightly over time. You should download it and use it locally for your
	// particular user interface needs.
	IconSVGURI string `json:"icon_svg_uri"`

	// SearchURI is a Scryfall API URI that you can request to begin
	// paginating over the cards in this set.
	SearchURI string `json:"search_uri"`
}

// ListSets lists all of the sets on Scryfall.
// TODO(serenst): Handle pagination.
func (c *Client) ListSets(ctx context.Context) ([]Set, error) {
	setsURL := fmt.Sprintf("%s/sets", baseURL)
	listResponse := &ListResponse{}
	err := c.doGETReq(ctx, setsURL, listResponse)
	if err != nil {
		return nil, err
	}

	sets := []Set{}
	err = json.Unmarshal(listResponse.Data, &sets)
	if err != nil {
		return nil, err
	}

	return sets, nil
}

// GetSet returns a set with the given set code.
func (c *Client) GetSet(ctx context.Context, code string) (Set, error) {
	setURL := fmt.Sprintf("%s/sets/%s", baseURL, code)
	set := Set{}
	err := c.doGETReq(ctx, setURL, &set)
	if err != nil {
		return Set{}, err
	}

	return set, nil
}