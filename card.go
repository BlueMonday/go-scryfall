package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
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

	// Name is the name of this card. If this card has multiple faces, this
	// field will contain both names separated by ␣//␣.
	Name string `json:"name"`

	// Layout is a computer-readable designation for this card’s
	// layout. See the layout article.
	Layout string `json:"layout"`

	// CMC is the card’s converted mana cost. Note that some funny cards
	// have fractional mana costs.
	CMC float64 `json:"cmc"`

	// TypeLine is the type line of this card.
	TypeLine string `json:"type_line"`

	// OracleText is the Oracle text for this card, if any.
	OracleText string `json:"oracle_text"`

	// ManaCost is the mana cost for this card. This value will be any
	// empty string "" if the cost is absent. Remember that per the game
	// rules, a missing mana cost and a mana cost of {0} are different values.
	ManaCost string `json:"mana_cost"`

	// Power is this card’s power, if any. Note that some cards have powers
	// that are not numeric, such as *.
	Power *string `json:"power"`

	// Toughness is this card’s toughness, if any. Note that some cards
	// have toughnesses that are not numeric, such as *.
	Toughness *string `json:"toughness"`

	// Loyalty is this loyalty if any. Note that some cards have loyalties
	// that are not numeric, such as X.
	Loyalty *string `json:"loyalty"`

	// LifeModifier is this card’s life modifier, if it is Vanguard
	// card. This value will contain a delta, such as +2.
	LifeModifier *string `json:"life_modifier"`

	// HandModifier is this card’s hand modifier, if it is Vanguard
	// card. This value will contain a delta, such as -1.
	HandModifier *string `json:"hand_modifier"`

	// Colors is this card’s colors.
	Colors []Color `json:"colors"`

	// ColorIndicator is the colors in this card’s color indicator, if
	// any. A nil value for this field indicates the card does not have one.
	ColorIndicator []Color `json:"color_indicator"`

	// ColorIdentity is this card’s color identity.
	ColorIdentity []Color `json:"color_identity"`

	// AllParts is a list of closely related cards, if any.
	AllParts []RelatedCard `json:"all_parts"`

	// CardFaces is An array of card Face objects, if this card is
	// multifaced.
	CardFaces []CardFace `json:"card_faces"`

	// Legalities is an object describing the legality of this card.
	Legalities Legalities `json:"legalities"`

	// Reserved is true if this card is on the Reserved List.
	Reserved bool `json:"reserved"`

	// EDHRECRank is this card’s overall rank/popularity on EDHREC. Not all
	// cards are ranked.
	EDHRECRank *int `json:"edhrec_rank"`

	// Set is this card’s set code.
	Set string `json:"set"`

	// SetName is this card’s full set name.
	SetName string `json:"set_name"`

	// CollectorNumber is this card’s collector number. Note that collector
	// numbers can contain non-numeric characters, such as letters or ★.
	CollectorNumber string `json:"collector_number"`

	// SetSearchURI is a link to where you can begin paginating this card’s
	// set on the Scryfall API.
	SetSearchURI string `json:"set_search_uri"`

	// ScryfallSetURI is a link to this card’s set on Scryfall’s website.
	ScryfallSetURI string `json:"scryfall_set_uri"`

	// ImageURIs is an object listing available imagery for this card.
	ImageURIs *ImageURIs `json:"image_uris"`

	// HighresImage is true if this card’s imagery is high resolution.
	HighresImage bool `json:"highres_image"`

	// Reprint is true if this card is a reprint.
	Reprint bool `json:"reprint"`

	// Digital is true if this is a digital card on Magic Online.
	Digital bool `json:"digital"`

	// Rarity is this card’s rarity. One of common, uncommon, rare, or
	// mythic.
	Rarity string `json:"rarity"`

	// FlavorText is the flavor text, if any.
	FlavorText *string `json:"flavor_text"`

	// Artist is the name of the illustrator of this card. Newly spoiled
	// cards may not have this field yet.
	Artist *string `json:"artist"`

	// IllustrationID is a unique identifier for the card artwork that
	// remains consistent across reprints. Newly spoiled cards may not have
	// this field yet.
	IllustrationID *string `json:"illustration_id"`

	// Frame is this card’s frame layout.
	Frame Frame `json:"frame"`

	// FullArt is true if this card’s artwork is larger than normal.
	FullArt bool `json:"full_art"`

	// Watermark is this card’s watermark, if any.
	Watermark *string `json:"watermark"`

	// BorderColor is this card’s border color: black, borderless, gold,
	// silver, or white.
	BorderColor string `json:"border_color"`

	// StorySpotlightNumber is this card’s story spotlight number, if any.
	StorySpotlightNumber *int `json:"story_spotlight_number"`

	// StorySpotlightURI is a URL to this cards’s story article, if any.
	StorySpotlightURI *string `json:"story_spotlight_uri"`

	// Timeshifted is true if this card is timeshifted.
	Timeshifted bool `json:"timeshifted"`

	// Colorshifted is true if this card is colorshifted.
	Colorshifted bool `json:"colorshifted"`

	// Futureshifted is true if this card is from the future.
	Futureshifted bool `json:"futureshifted"`
}

// RelatedCard is a card that is closely related to another card (because it
// calls it by name, or generates a token, or meld, etc).
type RelatedCard struct {
	// ID is a unique ID for this card in Scryfall’s database.
	ID string `json:"id"`

	// Name is the name of this particular related card.
	Name string `json:"name"`

	// URI is a URI where you can retrieve a full object describing this
	// card on Scryfall’s API.
	URI string `json:"uri"`
}

// CardFace is a face of a multifaced card.
type CardFace struct {
	// Name is the name of this particular face.
	Name string `json:"name"`

	// TypeLine is the type line of this particular face.
	TypeLine string `json:"type_line"`

	// OracleText is the Oracle text for this face, if any.
	OracleText *string `json:"oracle_text"`

	// ManaCost is the mana cost for this face. This value will be any
	// empty string "" if the cost is absent. Remember that per the game
	// rules, a missing mana cost and a mana cost of {0} are different values.
	ManaCost string `json:"mana_cost"`

	// Colors is this face’s colors.
	Colors []Color `json:"colors"`

	// ColorIndicator is the colors in this face’s color indicator, if any.
	ColorIndicator []Color `json:"color_indicator"`

	// Power is this face’s power, if any. Note that some cards have powers
	// that are not numeric, such as *.
	Power *string `json:"power"`

	// Toughness is this face’s toughness, if any.
	Toughness *string `json:"toughness"`

	// Loyalty is this face’s loyalty, if any.
	Loyalty *string `json:"loyalty"`

	// FlavorText is the flavor text printed on this face, if any.
	FlavorText *string `json:"flavor_text"`

	// IllustrationID is a unique identifier for the card face artwork that
	// remains consistent across reprints. Newly spoiled cards may not have
	// this field yet.
	IllustrationID *string `json:"illustration_id"`

	// ImageURIs is an object providing URIs to imagery for this face, if
	// this is a double-sided card. If this card is not double-sided, then the
	// image_uris property will be part of the parent object instead.
	ImageURIs ImageURIs `json:"image_uris"`
}

type ImageURIs struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	PNG        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Legality string

const (
	LegalityLegal      Legality = "legal"
	LegalityNotLegal   Legality = "not_legal"
	LegalityBanned     Legality = "banned"
	LegalityRestricted Legality = "restricted"
)

// Legalities describes the legality of a card across formats.
type Legalities struct {
	Standard     Legality `json:"standard"`
	Fronteir     Legality `json:"fronteir"`
	Modern       Legality `json:"modern"`
	Pauper       Legality `json:"pauper"`
	Legacy       Legality `json:"legacy"`
	Penny        Legality `json:"penny"`
	Vintage      Legality `json:"vintage"`
	Duel         Legality `json:"duel"`
	Commander    Legality `json:"commander"`
	OneVersusOne Legality `json:"1v1"`
	Future       Legality `json:"future"`
}

// ListCards lists all the cards in Scryfall's database.
// TODO(serenst): Handle pagination.
func (c *Client) ListCards(ctx context.Context) ([]Card, error) {
	cardsURL := fmt.Sprintf("%s/cards", baseURL)
	listResponse := &ListResponse{}
	err := c.doGETReq(ctx, cardsURL, listResponse)
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

	catalog := Catalog{}
	err = c.doGETReq(ctx, autocompleteCardURL.String(), &catalog)
	if err != nil {
		return Catalog{}, err
	}

	return catalog, nil
}

// GetRandomCard returns a random card.
func (c *Client) GetRandomCard(ctx context.Context) (Card, error) {
	randomCardURL := fmt.Sprintf("%s/cards/random", baseURL)
	card := Card{}
	err := c.doGETReq(ctx, randomCardURL, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// GetCardBySetCodeAndCollectorNumber returns a single card with the given
// set code and collector number.
func (c *Client) GetCardBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber string) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/%s/%s", baseURL, setCode, collectorNumber)
	card := Card{}
	err := c.doGETReq(ctx, cardURL, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// GetCard returns a single card with the given Scryfall ID.
func (c *Client) GetCard(ctx context.Context, id string) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/%s", baseURL, id)
	card := Card{}
	err := c.doGETReq(ctx, cardURL, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}
