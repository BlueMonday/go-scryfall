package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
)

// Ruling represents an Oracle ruling, Wizards of the Coast set release notes,
// or Scryfall notes for a particular card.
//
// If two cards have the same name, they will have the same set of rulings
// objects. If a card has rulings, it usually has more than one.
//
// Rulings with a scryfall source have been added by the Scryfall team, either
// to provide additional context for the card, or explain how the card works in an
// unofficial format (such as Duel Commander).
type Ruling struct {
	Source      string `json:"source"`
	PublishedAt string `json:"published_at"`
	Comment     string `json:"comment"`
}

func (c *Client) getRulings(ctx context.Context, url string) ([]Ruling, error) {
	listResponse := &ListResponse{}
	err := c.doGETReq(ctx, url, listResponse)
	if err != nil {
		return nil, err
	}

	rulings := []Ruling{}
	err = json.Unmarshal(listResponse.Data, &rulings)
	if err != nil {
		return nil, err
	}

	return rulings, nil
}

// GetRulingsByMTGOID returns rulings for a card with the given MTGO ID (also
// known as the Catalog ID). The ID can either be the card’s MTGO ID or its
// MTGO foil ID.
func (c *Client) GetRulingsByMTGOID(ctx context.Context, mtgoID int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("%s/cards/mtgo/%d/rulings", baseURL, mtgoID)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulingsBySetCodeAndCollectorNumber returns a list of rulings for the card
// with the given set code and collector number.
func (c *Client) GetRulingsBySetCodeAndCollectorNumber(ctx context.Context, setCode int, collectorNumber int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("%s/cards/%d/%d/rulings", baseURL, setCode, collectorNumber)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulingsByMultiverseID returns the rulings for a card with the given
// multiverse ID. If the card has multiple multiverse IDs,
// GetRulingsByMultiverseID can find either of them.
func (c *Client) GetRulingsByMultiverseID(ctx context.Context, multiverseID int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("%s/cards/multiverse/%d/rulings", baseURL, multiverseID)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulings returns the rulings for a card with the given Scryfall ID.
func (c *Client) GetRulings(ctx context.Context, id string) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("%s/cards/%s/rulings", baseURL, id)
	return c.getRulings(ctx, rulingsURL)
}
