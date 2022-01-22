package scryfall

import (
	"context"
	"fmt"
)

// Source indicates which company produced the ruling.
type Source string

const (
	// SourceWOTC is a Wizards of the Coast ruling source.
	SourceWOTC Source = "wotc"

	// SourceScryfall is a Scryfall ruling source.
	SourceScryfall Source = "scryfall"
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
	// Source indicates which company produced the ruling.
	Source Source `json:"source"`

	// PublishedAt is the date when the ruling or note was published.
	PublishedAt Date `json:"published_at"`

	// Comment is the text of the ruling.
	Comment string `json:"comment"`
}

func (c *Client) getRulings(ctx context.Context, url string) ([]Ruling, error) {
	rulings := []Ruling{}
	err := c.listGet(ctx, url, &rulings)
	if err != nil {
		return nil, err
	}

	return rulings, nil
}

// GetRulingsByMultiverseID returns the rulings for a card with the given
// multiverse ID. If the card has multiple multiverse IDs,
// GetRulingsByMultiverseID can find either of them.
func (c *Client) GetRulingsByMultiverseID(ctx context.Context, multiverseID int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("cards/multiverse/%d/rulings", multiverseID)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulingsByMTGOID returns rulings for a card with the given MTGO ID (also
// known as the Catalog ID). The ID can either be the card's MTGO ID or its
// MTGO foil ID.
func (c *Client) GetRulingsByMTGOID(ctx context.Context, mtgoID int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("cards/mtgo/%d/rulings", mtgoID)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulingsByArenaID returns rulings for a card with the given Magic: The
// Gathering Arena ID.
func (c *Client) GetRulingsByArenaID(ctx context.Context, arenaID int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("cards/arena/%d/rulings", arenaID)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulingsBySetCodeAndCollectorNumber returns a list of rulings for the card
// with the given set code and collector number.
func (c *Client) GetRulingsBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber int) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("cards/%s/%d/rulings", setCode, collectorNumber)
	return c.getRulings(ctx, rulingsURL)
}

// GetRulings returns the rulings for a card with the given Scryfall ID.
func (c *Client) GetRulings(ctx context.Context, id string) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("cards/%s/rulings", id)
	return c.getRulings(ctx, rulingsURL)
}
