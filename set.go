package scryfall

import (
	"context"
	"fmt"
)

// SetType is a computer-readable classification for a set.
type SetType string

const (
	// SetTypeCore is a yearly Magic core set (Tenth Edition, etc).
	SetTypeCore SetType = "Core"

	// SetTypeExpansion is a rotational expansion set in a block (Zendikar,
	// etc).
	SetTypeExpansion SetType = "expansion"

	// SetTypeMasters is a reprint set that contains no new cards (Modern
	// Masters, etc).
	SetTypeMasters SetType = "masters"

	// SetTypeMasterpiece is a set that contains masterpiece series premium
	// foil cards.
	SetTypeMasterpiece SetType = "masterpiece"

	// SetTypeFromTheVault is a From the Vault gift set.
	SetTypeFromTheVault SetType = "from_the_vault"

	// SetTypeSpellbook is a Spellbook series gift set.
	SetTypeSpellbook SetType = "spellbook"

	// SetTypePremiumDeck is a premium Deck Series decks set.
	SetTypePremiumDeck SetType = "premium_deck"

	// SetTypeDuelDeck is a Duel Decks set.
	SetTypeDuelDeck SetType = "duel_deck"

	// SetTypeDraftInnovation is a special draft set, like Conspiracy and Battlebond
	SetTypeDraftInnovation SetType = "draft_innovation"

	// SetTypeTreasureChest is a Magic Online treasure chest prize set.
	SetTypeTreasureChest SetType = "treasure_chest"

	// SetTypeCommander is a commander preconstructed set.
	SetTypeCommander SetType = "commander"

	// SetTypePlanechase is a Planechase set.
	SetTypePlanechase SetType = "planechase"

	// SetTypeArchenemy is an Archenemy set.
	SetTypeArchenemy SetType = "archenemy"

	// SetTypeVanguard is a Vanguard card set.
	SetTypeVanguard SetType = "vanguard"

	// SetTypeFunny is a funny un-set or set with funny promos (Unglued,
	// Happy Holidays, etc).
	SetTypeFunny SetType = "funny"

	// SetTypeStarter is a starter/introductory set (Portal, etc).
	SetTypeStarter SetType = "starter"

	// SetTypeBox is a gift box set.
	SetTypeBox SetType = "box"

	// SetTypePromo is a set that contains purely promotional cards.
	SetTypePromo SetType = "promo"

	// SetTypeToken is a set made up of tokens and emblems.
	SetTypeToken SetType = "token"

	// SetTypeMemorabilia is a set made up of gold-bordered, oversize, or
	// trophy cards that are not legal.
	SetTypeMemorabilia SetType = "memorabilia"
)

// Set is an object which represents a group of related Magic cards. All Card
// objects on Scryfall belong to exactly one set.
type Set struct {
	// ID is a unique ID for this set in Scryfall's database.
	ID string `json:"id"`

	// Code is the unique three or four-letter code for this set.
	Code string `json:"code"`

	// MTGOCode is the unique code for this set on MTGO, which may differ
	// from the regular code.
	MTGOCode *string `json:"mtgo_code"`

	// ArenaCode is the unique code for this set on Magic: The Gathering Arena,
	// which may differ from the regular code.
	ArenaCode *string `json:"arena_code"`

	// TCGplayerID is the set ID on TCGplayer's API, also known as the groupId.
	TCGplayerID *int `json:"tcgplayer_id"`

	// Name is the English name of the set.
	Name string `json:"name"`

	// URI is a link to this set object on Scryfall's API.
	URI string `json:"uri"`

	// ScryfallURI is a link to this card's permapage on Scryfall's website.
	ScryfallURI string `json:"scryfall_uri"`

	// SetType is a computer-readable classification for this set.
	SetType SetType `json:"set_type"`

	// ReleasedAt is the date the set was released (in GMT-8 Pacific
	// time). Not all sets have a known release date.
	ReleasedAt *Date `json:"released_at"`

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

	// FoilOnly is true if this set contains only foil cards.
	FoilOnly bool `json:"foil_only"`

	// IconSVGURI is a URI to an SVG file for this set's icon on Scryfall's
	// CDN. Hotlinking this image isn't recommended, because it may change
	// slightly over time. You should download it and use it locally for your
	// particular user interface needs.
	IconSVGURI string `json:"icon_svg_uri"`

	// SearchURI is a Scryfall API URI that you can request to begin
	// paginating over the cards in this set.
	SearchURI string `json:"search_uri"`
}

// ListSets lists all of the sets on Scryfall.
func (c *Client) ListSets(ctx context.Context) ([]Set, error) {
	sets := []Set{}
	err := c.listGet(ctx, "sets", &sets)
	if err != nil {
		return nil, err
	}

	return sets, nil
}

// GetSet returns a set with the given set code.
func (c *Client) GetSet(ctx context.Context, code string) (Set, error) {
	setURL := fmt.Sprintf("sets/%s", code)
	set := Set{}
	err := c.get(ctx, setURL, &set)
	if err != nil {
		return Set{}, err
	}

	return set, nil
}
