package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const duskDawnJSON = `{"object": "card", "id": "937dbc51-b589-4237-9fce-ea5c757f7c48", "oracle_id": "7bc3f92f-68a2-4934-afc4-89f6d0e8cf98", "multiverse_ids": [426912], "mtgo_id": 64026, "mtgo_foil_id": 64027, "name": "Dusk // Dawn", "uri": "https://api.scryfall.com/cards/akh/210", "scryfall_uri": "https://scryfall.com/card/akh/210?utm_source=api", "layout": "split", "highres_image": true, "image_uris": {"small": "https://img.scryfall.com/cards/small/en/akh/210a.jpg?1519586027", "normal": "https://img.scryfall.com/cards/normal/en/akh/210a.jpg?1519586027", "large": "https://img.scryfall.com/cards/large/en/akh/210a.jpg?1519586027", "png": "https://img.scryfall.com/cards/png/en/akh/210a.png?1519586027", "art_crop": "https://img.scryfall.com/cards/art_crop/en/akh/210a.jpg?1519586027", "border_crop": "https://img.scryfall.com/cards/border_crop/en/akh/210a.jpg?1519586027"}, "cmc": 9, "mana_cost": "{2}{W}{W} // {3}{W}{W}", "colors": ["W"], "color_identity": ["W"], "card_faces": [{"object": "card_face", "name": "Dusk", "mana_cost": "{2}{W}{W}", "type_line": "Sorcery", "oracle_text": "Destroy all creatures with power 3 or greater.", "illustration_id": "f3d63aed-2784-4ef5-9676-846b1e65e040"}, {"object": "card_face", "name": "Dawn", "mana_cost": "{3}{W}{W}", "type_line": "Sorcery", "oracle_text": "Aftermath (Cast this spell only from your graveyard. Then exile it.)\nReturn all creature cards with power 2 or less from your graveyard to your hand."}], "legalities": {"standard": "legal", "frontier": "legal", "modern": "legal", "pauper": "not_legal", "legacy": "legal", "penny": "not_legal", "vintage": "legal", "duel": "legal", "commander": "legal", "1v1": "legal", "future": "legal"}, "reserved": false, "reprint": false, "set": "akh", "set_name": "Amonkhet", "set_uri": "https://api.scryfall.com/sets/akh", "set_search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Aakh&unique=prints", "scryfall_set_uri": "https://scryfall.com/sets/akh?utm_source=api", "rulings_uri": "https://api.scryfall.com/cards/akh/210/rulings", "prints_search_uri": "https://api.scryfall.com/cards/search?order=set&q=%21%E2%80%9CDusk+%2F%2F+Dawn%E2%80%9D&unique=prints", "collector_number": "210", "digital": false, "rarity": "rare", "illustration_id": "f3d63aed-2784-4ef5-9676-846b1e65e040", "artist": "Noah Bradley", "frame": "2015", "frame_effect": "", "full_art": false, "border_color": "black", "timeshifted": false, "colorshifted": false, "futureshifted": false, "edhrec_rank": 942, "usd": "0.99", "tix": "0.15", "eur": "0.79", "related_uris": {"gatherer": "http://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=426912", "tcgplayer_decks": "http://decks.tcgplayer.com/magic/deck/search?contains=Dusk+%2F%2F+Dawn&page=1&partner=Scryfall", "edhrec": "http://edhrec.com/route/?cc=Dusk", "mtgtop8": "http://mtgtop8.com/search?MD_check=1&SB_check=1&cards=Dusk+%2F%2F+Dawn"}, "purchase_uris": {"tcgplayer": "https://scryfall.com/s/tcgplayer/129823", "cardmarket": "https://scryfall.com/s/mcm/296759", "cardhoarder": "https://www.cardhoarder.com/cards/64026?affiliate_id=scryfall&ref=card-profile&utm_campaign=affiliate&utm_medium=card&utm_source=scryfall"}}`

var duskDawn = Card{
	ID:            "937dbc51-b589-4237-9fce-ea5c757f7c48",
	OracleID:      "7bc3f92f-68a2-4934-afc4-89f6d0e8cf98",
	MultiverseIDs: []int{426912},
	MTGOID:        intPointer(64026),
	MTGOFoilID:    intPointer(64027),
	Name:          "Dusk // Dawn",
	URI:           "https://api.scryfall.com/cards/akh/210",
	ScryfallURI:   "https://scryfall.com/card/akh/210?utm_source=api",
	Layout:        LayoutSplit,
	HighresImage:  true,
	ImageURIs: &ImageURIs{
		Small:      "https://img.scryfall.com/cards/small/en/akh/210a.jpg?1519586027",
		Normal:     "https://img.scryfall.com/cards/normal/en/akh/210a.jpg?1519586027",
		Large:      "https://img.scryfall.com/cards/large/en/akh/210a.jpg?1519586027",
		PNG:        "https://img.scryfall.com/cards/png/en/akh/210a.png?1519586027",
		ArtCrop:    "https://img.scryfall.com/cards/art_crop/en/akh/210a.jpg?1519586027",
		BorderCrop: "https://img.scryfall.com/cards/border_crop/en/akh/210a.jpg?1519586027",
	},
	CMC:           9,
	ManaCost:      "{2}{W}{W} // {3}{W}{W}",
	Colors:        []Color{ColorWhite},
	ColorIdentity: []Color{ColorWhite},
	CardFaces: []CardFace{
		{
			Name:           "Dusk",
			ManaCost:       "{2}{W}{W}",
			TypeLine:       "Sorcery",
			OracleText:     stringPointer("Destroy all creatures with power 3 or greater."),
			IllustrationID: stringPointer("f3d63aed-2784-4ef5-9676-846b1e65e040"),
		},
		{
			Name:       "Dawn",
			ManaCost:   "{3}{W}{W}",
			TypeLine:   "Sorcery",
			OracleText: stringPointer("Aftermath (Cast this spell only from your graveyard. Then exile it.)\nReturn all creature cards with power 2 or less from your graveyard to your hand."),
		},
	},
	Legalities: Legalities{
		Standard:     "legal",
		Frontier:     "legal",
		Modern:       "legal",
		Pauper:       "not_legal",
		Legacy:       "legal",
		Penny:        "not_legal",
		Vintage:      "legal",
		Duel:         "legal",
		Commander:    "legal",
		OneVersusOne: "legal",
		Future:       "legal",
	},
	Reserved:        false,
	Reprint:         false,
	Set:             "akh",
	SetName:         "Amonkhet",
	SetURI:          "https://api.scryfall.com/sets/akh",
	SetSearchURI:    "https://api.scryfall.com/cards/search?order=set&q=e%3Aakh&unique=prints",
	ScryfallSetURI:  "https://scryfall.com/sets/akh?utm_source=api",
	RulingsURI:      "https://api.scryfall.com/cards/akh/210/rulings",
	PrintsSearchURI: "https://api.scryfall.com/cards/search?order=set&q=%21%E2%80%9CDusk+%2F%2F+Dawn%E2%80%9D&unique=prints",
	CollectorNumber: "210",
	Digital:         false,
	Rarity:          "rare",
	IllustrationID:  stringPointer("f3d63aed-2784-4ef5-9676-846b1e65e040"),
	Artist:          stringPointer("Noah Bradley"),
	Frame:           Frame2015,
	FrameEffect:     "",
	FullArt:         false,
	BorderColor:     "black",
	Timeshifted:     false,
	Colorshifted:    false,
	Futureshifted:   false,
	EDHRECRank:      intPointer(942),
	USD:             "0.99",
	Tix:             "0.15",
	EUR:             "0.79",
	RelatedURIs: RelatedURIs{
		Gatherer:       "http://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=426912",
		TCGPlayerDecks: "http://decks.tcgplayer.com/magic/deck/search?contains=Dusk+%2F%2F+Dawn&page=1&partner=Scryfall",
		EDHREC:         "http://edhrec.com/route/?cc=Dusk",
		MTGTop8:        "http://mtgtop8.com/search?MD_check=1&SB_check=1&cards=Dusk+%2F%2F+Dawn",
	},
	PurchaseURIs: PurchaseURIs{
		TCGPlayer:   "https://scryfall.com/s/tcgplayer/129823",
		CardMarket:  "https://scryfall.com/s/mcm/296759",
		CardHoarder: "https://www.cardhoarder.com/cards/64026?affiliate_id=scryfall&ref=card-profile&utm_campaign=affiliate&utm_medium=card&utm_source=scryfall",
	},
}

func TestListCards(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		if page != "2" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, `{"object": "list", "total_cards": 1000, "has_more": true, "next_page": "https://api.scryfall.com/cards?page=3", "data": [`+duskDawnJSON+`]}`)
	})
	client, ts, err := setupTestServer("/cards", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	opts := ListCardsOptions{
		Page: 2,
	}
	cards, err := client.ListCards(ctx, opts)
	if err != nil {
		t.Fatalf("Error listing cards: %v", err)
	}

	want := CardListResponse{
		Cards:      []Card{duskDawn},
		HasMore:    true,
		NextPage:   stringPointer("https://api.scryfall.com/cards?page=3"),
		TotalCards: 1000,
	}
	if !reflect.DeepEqual(cards, want) {
		t.Errorf("got: %#v want: %#v", cards, want)
	}
}
func TestSearchCards(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		q := query.Get("q")
		unique := query.Get("unique")
		order := query.Get("order")
		dir := query.Get("dir")
		includeExtras := query.Get("include_extras")
		page := query.Get("page")
		if q != "dusk" && unique != "cards" && order != "power" && dir != "auto" && includeExtras != "true" && page != "2" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, `{"object": "list", "total_cards": 1000, "has_more": true, "next_page": "https://api.scryfall.com/cards?page=3", "data": [`+duskDawnJSON+`]}`)
	})
	client, ts, err := setupTestServer("/cards/search", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	opts := SearchCardsOptions{
		Unique:        UniqueModeCards,
		Order:         OrderPower,
		Dir:           DirAuto,
		IncludeExtras: true,
		Page:          2,
	}
	cards, err := client.SearchCards(ctx, "dusk", opts)
	if err != nil {
		t.Fatalf("Error listing cards: %v", err)
	}

	want := CardListResponse{
		Cards:      []Card{duskDawn},
		HasMore:    true,
		NextPage:   stringPointer("https://api.scryfall.com/cards?page=3"),
		TotalCards: 1000,
	}
	if !reflect.DeepEqual(cards, want) {
		t.Errorf("got: %#v want: %#v", cards, want)
	}
}

func TestGetCardByName(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		set := query.Get("set")
		fuzzyName := query.Get("fuzzy")
		if set != "akh" && fuzzyName != "Dusk" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/named", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	opts := GetCardByNameOptions{
		Set: "akh",
	}
	card, err := client.GetCardByName(ctx, "Dusk", false, opts)
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestAutocompleteCard(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q != "thal" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, `{"object": "catalog", "total_items": 20, "data": ["Thallid", "Thorn Thallid", "Thalakos Seer", "Thalakos Scout", "Thalia's Lancers", "Thalakos Sentry", "Thallid Devourer", "Thalakos Deceiver", "Thalakos Drifters", "Thalakos Lowlands", "Thalakos Mistfolk", "Thallid Soothsayer", "Thallid Germinator", "Thalia's Lieutenant", "Thallid Shell-Dweller", "Thalia, Heretic Cathar", "Thalakos Dreamsower", "Thalia, Guardian of Thraben", "Tukatongue Thallid", "Lethal Sting"]}`)
	})
	client, ts, err := setupTestServer("/cards/autocomplete", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	autocompletions, err := client.AutocompleteCard(ctx, "thal")
	if err != nil {
		t.Fatalf("Error auto completing card: %v", err)
	}

	want := []string{
		"Thallid",
		"Thorn Thallid",
		"Thalakos Seer",
		"Thalakos Scout",
		"Thalia's Lancers",
		"Thalakos Sentry",
		"Thallid Devourer",
		"Thalakos Deceiver",
		"Thalakos Drifters",
		"Thalakos Lowlands",
		"Thalakos Mistfolk",
		"Thallid Soothsayer",
		"Thallid Germinator",
		"Thalia's Lieutenant",
		"Thallid Shell-Dweller",
		"Thalia, Heretic Cathar",
		"Thalakos Dreamsower",
		"Thalia, Guardian of Thraben",
		"Tukatongue Thallid",
		"Lethal Sting",
	}
	if !reflect.DeepEqual(autocompletions, want) {
		t.Errorf("got: %#v want: %#v", autocompletions, want)
	}
}

func TestGetRandomCard(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/random", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetRandomCard(ctx)
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestGetCardByMultiverseID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/multiverse/426912", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetCardByMultiverseID(ctx, 426912)
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestGetCardByArenaID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/arena/67330", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetCardByArenaID(ctx, 67330)
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestGetCardByMTGOID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/mtgo/64026", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetCardByMTGOID(ctx, 64026)
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestGetCardBySetCodeAndCollectorNumber(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/akh/210", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetCardBySetCodeAndCollectorNumber(ctx, "akh", "210")
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}

func TestGetCard(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, duskDawnJSON)
	})
	client, ts, err := setupTestServer("/cards/937dbc51-b589-4237-9fce-ea5c757f7c48", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	card, err := client.GetCard(ctx, "937dbc51-b589-4237-9fce-ea5c757f7c48")
	if err != nil {
		t.Fatalf("Error getting card: %v", err)
	}

	if !reflect.DeepEqual(card, duskDawn) {
		t.Errorf("got: %#v want: %#v", card, duskDawn)
	}
}
