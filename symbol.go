package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CardSymbol represents an illustrated symbol that may appear in card’s
// mana cost or Oracle text. Symbols are based on the notation used in the
// Comprehensive Rules.
//
// For more information about how the Scryfall API represents mana and costs, see
// the colors and costs overview: https://scryfall.com/docs/api/colors.
type CardSymbol struct {
	// Symbol is the plaintext symbol. Often surrounded with curly braces
	// {}. Note that not all symbols are ASCII text (for example, {∞}).
	Symbol string `json:"symbol"`

	// LooseVariant is an alternate version of this symbol, if it is
	// possible to write it without curly braces.
	LooseVariant *string `json:"loose_variant"`

	// English is an English snippet that describes this
	// symbol. Appropriate for use in alt text or other accessible
	// communication formats.
	English string `json:"english"`

	// Transposable is true if it is possible to write this symbol “backwards”. For
	// example, the official symbol {U/P} is sometimes written as {P/U} or {P\U} in
	// informal settings. Note that the Scryfall API never writes symbols backwards in
	// other responses. This field is provided for informational purposes.
	Transposable bool `json:"transposable"`

	// RepresentsMana is true if this is a mana symbol.
	RepresentsMana bool `json:"represents_mana"`

	// CMC is a decimal number representing this symbol’s converted mana
	// cost. Note that mana symbols from funny sets can have fractional
	// converted mana costs.
	CMC float64 `json:"cmc"`

	// AppearsInManaCosts is true if this symbol appears in a mana cost on
	// any Magic card. For example {20} has this field set to false because
	// {20} only appears in Oracle text, not mana costs.
	AppearsInManaCosts bool `json:"appears_in_mana_costs"`

	// Funny is true if this symbol is only used on funny cards or Un-cards.
	Funny bool `json:"funny"`

	// Color is an array of colors that this symbol represents.
	Colors []Color `json:"colors"`
}

// ListCardSymbols returns a list of all card symbols.
func (c *Client) ListCardSymbols(ctx context.Context) ([]CardSymbol, error) {
	cardsURL := fmt.Sprintf("%s/symbology", baseURL)
	req, err := http.NewRequest("GET", cardsURL, nil)
	if err != nil {
		return nil, err
	}

	listResponse := &ListResponse{}
	err = c.doReq(ctx, req, listResponse)
	if err != nil {
		return nil, err
	}

	symbols := []CardSymbol{}
	err = json.Unmarshal(listResponse.Data, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}
