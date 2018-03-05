package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

// GetRulings returns the rulings for a card with the given Scryfall ID.
func (c *Client) GetRulings(ctx context.Context, id string) ([]Ruling, error) {
	rulingsURL := fmt.Sprintf("%s/cards/%s/rulings", baseURL, id)
	req, err := http.NewRequest("GET", rulingsURL, nil)
	if err != nil {
		return nil, err
	}

	listResponse := &ListResponse{}
	err = c.doReq(ctx, req, listResponse)
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
