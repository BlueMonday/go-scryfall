package scryfall

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Layout categorizes the arrangement of card parts, faces, and other bounded
// regions on cards. The layout can be used to programmatically determine which
// other properties on a card you can expect.
type Layout string

const (
	LayoutNormal           Layout = "normal"
	LayoutSplit            Layout = "split"
	LayoutFlip             Layout = "flip"
	LayoutTransform        Layout = "transform"
	LayoutMeld             Layout = "meld"
	LayoutLeveler          Layout = "leveler"
	LayoutPlanar           Layout = "planar"
	LayoutScheme           Layout = "scheme"
	LayoutVanguard         Layout = "vanguard"
	LayoutToken            Layout = "token"
	LayoutDoubleFacedToken Layout = "double_faced_token"
	LayoutEmblem           Layout = "emblem"
	LayoutAugment          Layout = "augment"
	LayoutHost             Layout = "host"
)

// Legality is the legality of a card in a particular format.
type Legality string

const (
	LegalityLegal      Legality = "legal"
	LegalityNotLegal   Legality = "not_legal"
	LegalityBanned     Legality = "banned"
	LegalityRestricted Legality = "restricted"
)

// Frame tracks the major edition of the card frame of used for the re/print in
// question. The frame has gone though several major revisions in Magic’s
// lifetime.
type Frame string

const (
	Frame1993   Frame = "1993"
	Frame1997   Frame = "1997"
	Frame2003   Frame = "2003"
	Frame2015   Frame = "2015"
	FrameFuture Frame = "future"
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
	Amazon          string `json:"amazon"`
	Ebay            string `json:"ebay"`
	TCGPlayer       string `json:"tcgplayer"`
	MagicCardMarket string `json:"magiccardmarket"`
	CardHoarder     string `json:"cardhoarder"`
	CardKingdom     string `json:"card_kingdom"`
	MTGOTraders     string `json:"mtgo_traders"`
	CoolStuffInc    string `json:"coolstuffinc"`
}

// ListCards lists all the cards in Scryfall's database.
// TODO(serenst): Handle pagination.
func (c *Client) ListCards(ctx context.Context) ([]Card, error) {
	cardsURL := fmt.Sprintf("%s/cards", baseURL)
	cards := []Card{}
	err := c.doListGETReq(ctx, cardsURL, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
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

	// UniqueModePrint returns all prints for all cards matched (disables
	// rollup). For example, if your search matches more than one print of
	// Pacifism, all matching prints will be returned.
	UniqueModePrint UniqueMode = "print"
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

// SearchConfig holds the configuration used to search for a card.
type SearchConfig struct {
	// Query is the full text search query. See the search reference docs
	// for more information on the full text search query format:
	// https://scryfall.com/docs/reference.
	Query string `url:"q"`

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
}

// SearchCards returns a list cards found using a full text search.
func (c *Client) SearchCards(ctx context.Context, searchConfig SearchConfig) ([]Card, error) {
	cardsURL, err := url.Parse(fmt.Sprintf("%s/cards/search", baseURL))
	if err != nil {
		return nil, err
	}

	values, err := query.Values(searchConfig)
	if err != nil {
		return nil, err
	}
	cardsURL.RawQuery = values.Encode()

	cards := []Card{}
	err = c.doListGETReq(ctx, cardsURL.String(), &cards)
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

func (c *Client) getCard(ctx context.Context, url string) (Card, error) {
	card := Card{}
	err := c.doGETReq(ctx, url, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// GetRandomCard returns a random card.
func (c *Client) GetRandomCard(ctx context.Context) (Card, error) {
	randomCardURL := fmt.Sprintf("%s/cards/random", baseURL)
	return c.getCard(ctx, randomCardURL)
}

// GetCardByMultiverseID returns a single card with the given Multiverse ID. If
// the card has multiple multiverse IDs, GetCardByMultiverseID can find either of
// them.
func (c *Client) GetCardByMultiverseID(ctx context.Context, multiverseID int) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/multiverse/%d", baseURL, multiverseID)
	return c.getCard(ctx, cardURL)
}

// GetCardByMTGOID returns a single card with the given MTGO ID (also known as
// the Catalog ID). The ID can either be the card’s MTGO ID or its MTGO foil
// ID.
func (c *Client) GetCardByMTGOID(ctx context.Context, mtgoID int) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/mtgo/%d", baseURL, mtgoID)
	return c.getCard(ctx, cardURL)
}

// GetCardBySetCodeAndCollectorNumber returns a single card with the given
// set code and collector number.
func (c *Client) GetCardBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber string) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/%s/%s", baseURL, setCode, collectorNumber)
	return c.getCard(ctx, cardURL)
}

// GetCard returns a single card with the given Scryfall ID.
func (c *Client) GetCard(ctx context.Context, id string) (Card, error) {
	cardURL := fmt.Sprintf("%s/cards/%s", baseURL, id)
	return c.getCard(ctx, cardURL)
}
