package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Card represents individual Magic: The Gathering cards that players could
// obtain and add to their collection (with a few minor exceptions).
type Card struct {
	// ID is a unique ID for this card in Scryfall’s database.
	ID string `json:"id"`

	// OracleID is a unique ID for this card’s oracle identity. This value
	// is consistent across reprinted card editions, and unique among
	// different cards with the same name (tokens, Unstable variants, etc).
	OracleID string `json:"oracle_id"`

	// MultiverseIDs is this card’s multiverse IDs on Gatherer, if any, as
	// an array of integers. Note that Scryfall includes many promo cards,
	// tokens, and other esoteric objects that do not have these identifiers.
	MultiverseIDs []int `json:"multiverse_ids"`

	// MTGOID is this card’s Magic Online ID (also known as the Catalog
	// ID), if any. A large percentage of cards are not available on Magic
	// Online and do not have this ID.
	MTGOID *int `json:"mtgo_id"`

	// MTGOFoilID is this card’s foil Magic Online ID (also known as the
	// Catalog ID), if any. A large percentage of cards are not available on
	// Magic Online and do not have this ID.
	MTGOFoilID *int `json:"mtgo_foil_id"`

	// URI is a link to this card object on Scryfall’s API.
	URI string `json:"uri"`

	// ScryfallURI is a link to this card’s permapage on Scryfall’s website.
	ScryfallURI string `json:"scryfall_uri"`

	// PrintsSearchURI is a link to where you can begin paginating all
	// re/prints for this card on Scryfall’s API.
	PrintsSearchURI string `json:"prints_search_uri"`

	// RulingsURI is a link to this card’s rulings on Scryfall’s API.
	RulingsURI string `json:"rulings_uri"`

	// TODO(serenst): Add the remaining fields.
}

// ListCards lists all the cards in Scryfall's database.
// TODO(serenst): Handle pagination.
func (c *Client) ListCards(ctx context.Context) ([]Card, error) {
	cardsURL := fmt.Sprintf("%s/cards", baseURL)
	req, err := http.NewRequest("GET", cardsURL, nil)
	if err != nil {
		return nil, err
	}

	listResponse := &ListResponse{}
	err = c.doReq(ctx, req, listResponse)
	if err != nil {
		return nil, err
	}

	cards := []Card{}
	err = json.Unmarshal(listResponse.Data, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// AutocompleteCard returns a Catalog containing up to 20 full English card
// names that could be autocompletions of the given string parameter.
func (c *Client) AutocompleteCard(ctx context.Context, s string) (Catalog, error) {
	autocompleteCardURL, err := url.Parse(fmt.Sprintf("%s/cards/autocomplete", baseURL))
	if err != nil {
		return Catalog{}, err
	}

	values := autocompleteCardURL.Query()
	values.Set("q", s)
	autocompleteCardURL.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", autocompleteCardURL.String(), nil)
	if err != nil {
		return Catalog{}, err
	}

	catalog := Catalog{}
	err = c.doReq(ctx, req, &catalog)
	if err != nil {
		return Catalog{}, err
	}

	return catalog, nil
}

// GetRandomCard returns a random card.
func (c *Client) GetRandomCard(ctx context.Context) (Card, error) {
	randomCardURL := fmt.Sprintf("%s/cards/random", baseURL)
	req, err := http.NewRequest("GET", randomCardURL, nil)
	if err != nil {
		return Card{}, err
	}

	card := Card{}
	err = c.doReq(ctx, req, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// GetCard returns a single card with the given Scryfall ID.
func (c *Client) GetCard(ctx context.Context, id string) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/%s", baseURL, id)
	req, err := http.NewRequest("GET", cardURL, nil)
	if err != nil {
		return Card{}, err
	}

	card := Card{}
	err = c.doReq(ctx, req, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}
