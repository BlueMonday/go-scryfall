package scryfall

import (
	"context"
	"fmt"
	"net/url"
)

// CardSymbol represents an illustrated symbol that may appear in card's
// mana cost or Oracle text. Symbols are based on the notation used in the
// Comprehensive Rules.
//
// For more information about how the Scryfall API represents mana and costs, see
// the colors and costs overview: https://scryfall.com/docs/api/colors.
type CardSymbol struct {
	// Object is the content type for this object, always card_symbol.
	Object string `json:"object"`

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

	// CMC is a decimal number representing this symbol's converted mana
	// cost. Note that mana symbols from funny sets can have fractional
	// converted mana costs.
	//
	// Deprecated: Use ManaValue instead.
	CMC float64 `json:"cmc"`

	// ManaValue is a decimal number representing this symbol’s mana value
	// (also knowns as the converted mana cost). Note that mana symbols
	// from funny sets can have fractional mana values.
	ManaValue *float64 `json:"mana_value"`

	// AppearsInManaCosts is true if this symbol appears in a mana cost on
	// any Magic card. For example {20} has this field set to false because
	// {20} only appears in Oracle text, not mana costs.
	AppearsInManaCosts bool `json:"appears_in_mana_costs"`

	// Funny is true if this symbol is only used on funny cards or Un-cards.
	Funny bool `json:"funny"`

	// Color is an array of colors that this symbol represents.
	Colors []Color `json:"colors"`

	// Hybrid is true if the symbol is a hybrid mana symbol. Note that
	// monocolor Phyrexian symbols aren’t considered hybrid.
	Hybrid bool `json:"hybrid"`

	// Phyrexian is true if the symbol is a Phyrexian mana symbol, i.e. it
	// can be paid with 2 life.
	Phyrexian bool `json:"phyrexian"`

	// GathererAlternates is an array of plaintext versions of this symbol
	// that Gatherer uses on old cards to describe original printed
	// text. For example: {W} has ["oW", "ooW"] as alternates.
	GathererAlternates []string `json:"gatherer_alternates"`

	// SVGURI is a URI to an SVG image of this symbol on Scryfall’s CDNs.
	SVGURI *string `json:"svg_uri"`
}

// ManaCost is Scryfall's interpretation of a mana cost.
type ManaCost struct {
	// Cost is the normalized cost, with correctly-ordered and wrapped mana
	// symbols.
	Cost string `json:"cost"`

	// CMC is the converted mana cost. If you submit Un-set mana symbols,
	// this decimal could include fractional parts.
	CMC float64 `json:"cmc"`

	// Colors is the colors of the given cost.
	Colors []Color `json:"colors"`

	// Colorless is true if the cost is colorless.
	Colorless bool `json:"colorless"`

	// Monocolored is true if the cost is monocolored.
	Monocolored bool `json:"monocolored"`

	// Multicolored is true if the cost is multicolored.
	Multicolored bool `json:"multicolored"`
}

// ListCardSymbols returns a list of all card symbols.
func (c *Client) ListCardSymbols(ctx context.Context) ([]CardSymbol, error) {
	symbols := []CardSymbol{}
	err := c.listGet(ctx, "symbology", &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}

// ParseManaCost parses a string mana cost and returns Scryfall's interpretation.
//
// The server understands most community shorthand for mana costs (such as 2WW
// for {2}{W}{W}). Symbols can also be out of order, lowercase, or have multiple
// colorless costs (such as 2{g}2 for {4}{G}).
func (c *Client) ParseManaCost(ctx context.Context, cost string) (ManaCost, error) {
	values := url.Values{}
	values.Set("cost", cost)
	parseManaURL := fmt.Sprintf("symbology/parse-mana?%s", values.Encode())

	manaCost := ManaCost{}
	err := c.get(ctx, parseManaURL, &manaCost)
	if err != nil {
		return ManaCost{}, err
	}

	return manaCost, nil
}
