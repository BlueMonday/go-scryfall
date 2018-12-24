package scryfall

import (
	"context"
	"fmt"
	"net/url"

	qs "github.com/google/go-querystring/query"
)

// Layout categorizes the arrangement of card parts, faces, and other bounded
// regions on cards. The layout can be used to programmatically determine which
// other properties on a card you can expect.
type Layout string

const (
	// LayoutNormal is a standard Magic card layout with one face.
	LayoutNormal Layout = "normal"

	// LayoutSplit is a split-faced card layout.
	LayoutSplit Layout = "split"

	// LayoutFlip is a card layout that inverts vertically with the flip
	// keyword.
	LayoutFlip Layout = "flip"

	// LayoutTransform is a double-sided card layout that transforms.
	LayoutTransform Layout = "transform"

	// LayoutMeld is a card layout with meld parts printed on the back.
	LayoutMeld Layout = "meld"

	// LayoutLeveler is a level up card layout.
	LayoutLeveler Layout = "leveler"

	// LayoutSaga is saga card layout.
	LayoutSaga Layout = "saga"

	// LayoutPlanar is a plane and phenomenon card layout.
	LayoutPlanar Layout = "planar"

	// LayoutScheme is a scheme card layout.
	LayoutScheme Layout = "scheme"

	// LayoutVanguard is a vanguard card layout.
	LayoutVanguard Layout = "vanguard"

	// LayoutToken is a token card layout.
	LayoutToken Layout = "token"

	// LayoutDoubleFacedToken is a card token layout with another token
	// printed on the back.
	LayoutDoubleFacedToken Layout = "double_faced_token"

	// LayoutEmblem is an emblem card layout.
	LayoutEmblem Layout = "emblem"

	// LayoutAugment is an augment card layout.
	LayoutAugment Layout = "augment"

	// LayoutHost is host card layout.
	LayoutHost Layout = "host"
)

// Legality is the legality of a card in a particular format.
type Legality string

const (
	// LegalityLegal indicates the card is legal in the format.
	LegalityLegal Legality = "legal"

	// LegalityNotLegal indicates the card is not legal in the format.
	LegalityNotLegal Legality = "not_legal"

	// LegalityBanned indicates the card is banned in the format.
	LegalityBanned Legality = "banned"

	// LegalityRestricted indicates the card is restricted to one copy in
	// the format.
	LegalityRestricted Legality = "restricted"
)

// Frame tracks the major edition of the card frame of used for the re/print in
// question. The frame has gone though several major revisions in Magic’s
// lifetime.
type Frame string

const (
	// Frame1993 is the original Magic card frame, starting from Limited
	// Edition Alpha.
	Frame1993 Frame = "1993"

	// Frame1997 is the updated classic frame starting from Mirage block.
	Frame1997 Frame = "1997"

	// Frame2003 is the “modern” Magic card frame, introduced in Eighth
	// Edition and Mirrodin block.
	Frame2003 Frame = "2003"

	// Frame2015 is the holofoil-stamp Magic card frame, introduced in
	// Magic 2015.
	Frame2015 Frame = "2015"

	// FrameFuture is the frame used on cards from the future.
	FrameFuture Frame = "future"
)

// FrameEffect tracks additional frame artwork applied over a particular
// frame. For example, there are both 2003 and 2015-frame cards with the
// Nyx-touched effect.
type FrameEffect string

const (
	// FrameEffectLegendary is the legendary crown introduced in Dominaria.
	FrameEffectLegendary FrameEffect = "legendary"

	// FrameEffectMiracle is the miracle frame effect.
	FrameEffectMiracle FrameEffect = "miracle"

	// FrameEffectNyxTouched is the Nyx-touched frame effect.
	FrameEffectNyxTouched FrameEffect = "nyxtouched"

	// FrameEffectDraft is the draft-matters frame effect.
	FrameEffectDraft FrameEffect = "draft"

	// FrameEffectDevoid is the Devoid frame effect.
	FrameEffectDevoid FrameEffect = "devoid"

	// FrameEffectTombstone is the Odyssey tombstone mark frame effect.
	FrameEffectTombstone FrameEffect = "tombstone"

	// FrameEffectColorShifted is the colorshifted frame effect.
	FrameEffectColorShifted FrameEffect = "colorshifted"

	// FrameEffectSunMoonDFC is the sun and moon transform marks frame
	// effect.
	FrameEffectSunMoonDFC FrameEffect = "sunmoondfc"

	// FrameEffectCompassLandDFC is the compass and land transform marks
	// frame effect.
	FrameEffectCompassLandDFC FrameEffect = "compasslanddfc"

	// FrameEffectOriginPWDFC is the Origins and planeswalker transform
	// marks frame effect.
	FrameEffectOriginPWDFC FrameEffect = "originpwdfc"

	// FrameEffectMoonEldraziDFC is the moon and Eldrazi transform marks
	// frame effect.
	FrameEffectMoonEldraziDFC FrameEffect = "mooneldrazidfc"
)

// Card represents individual Magic: The Gathering cards that players could
// obtain and add to their collection (with a few minor exceptions).
type Card struct {
	// ArenaID is this card’s Arena ID, if any. A large percentage of cards
	// are not available on Arena and do not have this ID.
	ArenaID *int `json:"arena_id"`

	// ID is a unique ID for this card in Scryfall’s database.
	ID string `json:"id"`

	// Lang is the language code for this printing.
	Lang string `json:"lang"`

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
	Layout Layout `json:"layout"`

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

	// SetURI is a link to this card's set on Scryfall’s API.
	SetURI string `json:"set_uri"`

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

	// FrameEffect is this card's frame effect, if any.
	FrameEffect FrameEffect `json:"frame_effect"`

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

	// USD is the price of the card in US dollars.
	USD string `json:"usd"`

	// Tix is the price of the card in MTGO event tickets.
	Tix string `json:"tix"`

	// EUR is the price of the card in Euros.
	EUR string `json:"eur"`

	// RelatedURIs contains links related to a card.
	RelatedURIs RelatedURIs `json:"related_uris"`

	// PurchaseURIs contains links to the card on online card stores.
	PurchaseURIs PurchaseURIs `json:"purchase_uris"`
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

// ImageURIs contains links to the different image sizes and crops for a given
// card.
type ImageURIs struct {
	// Small is a small full card image. Designed for use as thumbnail or
	// list icon.
	Small string `json:"small"`

	// Normal is a medium-sized full card image.
	Normal string `json:"normal"`

	// Large is a large full card image.
	Large string `json:"large"`

	// PNG is a transparent, rounded full card PNG. This is the best image
	// to use for videos or other high-quality content.
	PNG string `json:"png"`

	// ArtCrop is a rectangular crop of the card’s art only. Not guaranteed
	// to be perfect for cards with outlier designs or strange frame
	// arrangements
	ArtCrop string `json:"art_crop"`

	// BorderCrop is a full card image with the rounded corners and the
	// majority of the border cropped off. Designed for dated contexts where
	// rounded images can’t be used.
	BorderCrop string `json:"border_crop"`
}

// Legalities describes the legality of a card across formats.
type Legalities struct {
	Standard     Legality `json:"standard"`
	Frontier     Legality `json:"frontier"`
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

// RelatedURIs contains links related to a card.
type RelatedURIs struct {
	Gatherer       string `json:"gatherer"`
	TCGPlayerDecks string `json:"tcgplayer_decks"`
	EDHREC         string `json:"edhrec"`
	MTGTop8        string `json:"mtgtop8"`
}

// PurchaseURIs contains links to the card on online card stores.
type PurchaseURIs struct {
	TCGPlayer   string `json:"tcgplayer"`
	CardMarket  string `json:"cardmarket"`
	CardHoarder string `json:"cardhoarder"`
}

// CardListResponse represents a requested sequence of card
// objects. CardListResponse objects may be paginated, and also include
// information about issues raised when generating the list.
type CardListResponse struct {
	// Cards is a list of the requested cards.
	Cards []Card `json:"data"`

	// HasMore is true if this List is paginated and there is a page beyond
	// the current page.
	HasMore bool `json:"has_more"`

	// NextPage contains a full API URI to next page if there is a page
	// beyond the current page.
	NextPage *string `json:"next_page"`

	// TotalCards contains the total number of cards found across all
	// pages.
	TotalCards int `json:"total_cards"`

	// Warnings is a list of human-readable warnings issued when generating
	// this list, as strings. Warnings are non-fatal issues that the API
	// discovered with your input. In general, they indicate that the List
	// will not contain the all of the information you requested. You should
	// fix the warnings and re-submit your request.
	Warnings []string `json:"warnings"`
}

// ListCardsOptions holds the options used to list cards.
type ListCardsOptions struct {
	// Page is the page number to return. Page numbers start at 1 and the
	// default is 1.
	Page int `url:"page,omitempty"`
}

// ListCards lists all the cards in Scryfall's database.
func (c *Client) ListCards(ctx context.Context, opts ListCardsOptions) (CardListResponse, error) {
	values, err := qs.Values(opts)
	if err != nil {
		return CardListResponse{}, err
	}

	cardsURL := fmt.Sprintf("cards?%s", values.Encode())
	result := CardListResponse{}
	err = c.get(ctx, cardsURL, &result)
	if err != nil {
		return CardListResponse{}, err
	}

	return result, nil
}

// UniqueMode specifies whether Scryfall should remove duplicates from search
// results.
type UniqueMode string

const (
	// UniqueModeCards removes duplicate gameplay objects (cards that share
	// a name and have the same functionality). For example, if your search
	// matches more than one print of Pacifism, only one copy of Pacifism will
	// be returned.
	UniqueModeCards UniqueMode = "cards"

	// UniqueModeArt returns only one copy of each unique artwork for
	// matching cards. For example, if your search matches more than one print
	// of Pacifism, one card with each different illustration for Pacifism
	// will be returned, but any cards that duplicate artwork already in the
	// results will be omitted.
	UniqueModeArt UniqueMode = "art"

	// UniqueModePrints returns all prints for all cards matched (disables
	// rollup). For example, if your search matches more than one print of
	// Pacifism, all matching prints will be returned.
	UniqueModePrints UniqueMode = "prints"
)

// Order is a method used to sort cards.
type Order string

const (
	// OrderName sorts cards by name, A → Z.
	OrderName Order = "name"

	// OrderSet sorts cards by their set and collector number: oldest →
	// newest.
	OrderSet Order = "set"

	// OrderRarity sorts cards by their rarity: Common → Mythic.
	OrderRarity Order = "rarity"

	// OrderColor sorts cards by their color and color identity: WUBRG →
	// multicolor → colorless.
	OrderColor Order = "color"

	// OrderUSD sorts cards by their lowest known U.S. Dollar price: 0.01 →
	// highest, null last.
	OrderUSD Order = "usd"

	// OrderTix sorts cards by their lowest known TIX price: 0.01 →
	// highest, null last.
	OrderTix Order = "tix"

	// OrderEUR sorts cards by their lowest known Euro price: 0.01 →
	// highest, null last.
	OrderEUR Order = "eur"

	// OrderCMC sorts cards by their converted mana cost: 0 → highest.
	OrderCMC Order = "cmc"

	// OrderPower sorts cards by their power: null → highest.
	OrderPower Order = "power"

	// OrderToughness sorts cards by their toughness: null → highest.
	OrderToughness Order = "toughness"

	// OrderEDHREC sorts cards by their EDHREC ranking: lowest → highest.
	OrderEDHREC Order = "edhrec"

	// OrderArtist sorts cards by their front-side artist name: A → Z.
	OrderArtist Order = "artist"
)

// Dir is a direction used to sort cards.
type Dir string

const (
	// DirAuto lets Scryfall automatically choose the most intuitive
	// direction to sort.
	DirAuto Dir = "auto"

	// DirAsc sorts cards in ascending order.
	DirAsc Dir = "asc"

	// DirDesc sorts cards in descending order.
	DirDesc Dir = "desc"
)

// SearchCardsOptions holds the options used to search for cards.
type SearchCardsOptions struct {
	// Unique is the strategy for omitting similar cards. The default
	// strategy is UniqueModeCards.
	Unique UniqueMode `url:"unique,omitempty"`

	// Order is the method used to sort the cards. The default method is
	// OrderName.
	Order Order `url:"order,omitempty"`

	// Dir is the direction to sort the cards. The default direction is
	// DirAuto.
	Dir Dir `url:"dir,omitempty"`

	// IncludeExtras determines whether extra cards (tokens, planes, etc.)
	// should be included.
	IncludeExtras bool `url:"include_extras,omitempty"`

	// Page is the page number to return. Page numbers start at 1 and the
	// default is 1.
	Page int `url:"page,omitempty"`
}

// SearchCards returns a list cards found using a full text search. The query
// parameter is the full text search query. See the search reference docs for more
// information on the full text search query format:
// https://scryfall.com/docs/reference.
func (c *Client) SearchCards(ctx context.Context, query string, opts SearchCardsOptions) (CardListResponse, error) {
	values, err := qs.Values(opts)
	if err != nil {
		return CardListResponse{}, err
	}
	values.Set("q", query)
	cardsURL := fmt.Sprintf("cards/search?%s", values.Encode())

	result := CardListResponse{}
	err = c.get(ctx, cardsURL, &result)
	if err != nil {
		return CardListResponse{}, err
	}

	return result, nil
}

func (c *Client) getCard(ctx context.Context, url string) (Card, error) {
	card := Card{}
	err := c.get(ctx, url, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// GetCardByNameOptions holds the options used to get a card by name.
type GetCardByNameOptions struct {
	// Set limits the search to the specified set.
	Set string `url:"set,omitempty"`
}

// GetCardByName returns a Card based on a name search string. This method is
// designed for building chat bots, forum bots, and other services that need card
// details quickly.
//
// If the exact parameter is set to true, a card with that exact name is
// returned. Otherwise, an error is returned because no card matches.
//
// If the exact parameter is set to false and a card name matches that string,
// then that card is returned. If not, a fuzzy search is executed for your card
// name. The server allows misspellings and partial words to be provided. For
// example: jac bel will match Jace Beleren.
//
// When fuzzy searching, a card is returned if the server is confident that you
// unambiguously identified a unique name with your string. Otherwise, you will
// receive an error describing the problem: either more than 1 one card matched
// your search, or zero cards matched.
//
// For both exact and fuzzy, card names are case-insensitive and punctuation is
// optional (you can drop apostrophes and periods etc). For example: fIReBALL is
// the same as Fireball and smugglers copter is the same as Smuggler's Copter
func (c *Client) GetCardByName(ctx context.Context, name string, exact bool, opts GetCardByNameOptions) (Card, error) {
	values, err := qs.Values(opts)
	if err != nil {
		return Card{}, err
	}

	if exact {
		values.Set("exact", name)
	} else {
		values.Set("fuzzy", name)
	}

	cardURL := fmt.Sprintf("cards/named?%s", values.Encode())
	return c.getCard(ctx, cardURL)
}

// AutocompleteCard returns a slice containing up to 20 full English card names
// that could be autocompletions of the given string parameter.
func (c *Client) AutocompleteCard(ctx context.Context, s string) ([]string, error) {
	values := url.Values{}
	values.Set("q", s)
	autocompleteCardURL := fmt.Sprintf("cards/autocomplete?%s", values.Encode())

	catalog := Catalog{}
	err := c.get(ctx, autocompleteCardURL, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog.Data, nil
}

// GetRandomCard returns a random card.
func (c *Client) GetRandomCard(ctx context.Context) (Card, error) {
	return c.getCard(ctx, "cards/random")
}

// GetCardBySetCodeAndCollectorNumber returns a single card with the given
// set code and collector number.
func (c *Client) GetCardBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber string) (Card, error) {
	cardURL := fmt.Sprintf("cards/%s/%s", setCode, collectorNumber)
	return c.getCard(ctx, cardURL)
}

// GetCardByMultiverseID returns a single card with the given Multiverse ID. If
// the card has multiple multiverse IDs, GetCardByMultiverseID can find either of
// them.
func (c *Client) GetCardByMultiverseID(ctx context.Context, multiverseID int) (Card, error) {
	cardURL := fmt.Sprintf("cards/multiverse/%d", multiverseID)
	return c.getCard(ctx, cardURL)
}

// GetCardByMTGOID returns a single card with the given MTGO ID (also known as
// the Catalog ID). The ID can either be the card’s MTGO ID or its MTGO foil
// ID.
func (c *Client) GetCardByMTGOID(ctx context.Context, mtgoID int) (Card, error) {
	cardURL := fmt.Sprintf("cards/mtgo/%d", mtgoID)
	return c.getCard(ctx, cardURL)
}

// GetCardByArenaID returns a single card with the given Magic: The Gathering
// Arena ID.
func (c *Client) GetCardByArenaID(ctx context.Context, arenaID int) (Card, error) {
	cardURL := fmt.Sprintf("cards/arena/%d", arenaID)
	return c.getCard(ctx, cardURL)
}

// GetCardByTCGPlayerID returns a single card with the given TCGPlayer ID, also
// known as the productId on TCGPlayer’s API.
func (c *Client) GetCardByTCGPlayerID(ctx context.Context, tcgPlayerID int) (Card, error) {
	cardURL := fmt.Sprintf("cards/tcgplayer/%d", tcgPlayerID)
	return c.getCard(ctx, cardURL)
}

// GetCard returns a single card with the given Scryfall ID.
func (c *Client) GetCard(ctx context.Context, id string) (Card, error) {
	cardURL := fmt.Sprintf("cards/%s", id)
	return c.getCard(ctx, cardURL)
}
