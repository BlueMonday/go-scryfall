package scryfall

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TODO: Replace urls with the new
// https://scryfall.com/blog/upcoming-api-changes-to-scryfall-image-uris-and-download-uris-221
const duskDawnJSON = `{"object":"card","id":"937dbc51-b589-4237-9fce-ea5c757f7c48","oracle_id":"7bc3f92f-68a2-4934-afc4-89f6d0e8cf98","multiverse_ids":[426912],"mtgo_id":64026,"tcgplayer_id":129823,"cardmarket_id":296759,"name":"Dusk // Dawn","lang":"en","released_at":"2017-04-28","uri":"https://api.scryfall.com/cards/937dbc51-b589-4237-9fce-ea5c757f7c48","scryfall_uri":"https://scryfall.com/card/akh/210/dusk-dawn?utm_source=api","layout":"split","highres_image":true,"image_status":"highres_scan","image_uris":{"small":"https://cards.scryfall.io/small/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330","normal":"https://cards.scryfall.io/normal/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330","large":"https://cards.scryfall.io/large/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330","png":"https://cards.scryfall.io/png/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.png?1549941330","art_crop":"https://cards.scryfall.io/art_crop/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330","border_crop":"https://cards.scryfall.io/border_crop/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330"},"mana_cost":"{2}{W}{W} // {3}{W}{W}","cmc":9.0,"type_line":"Sorcery // Sorcery","colors":["W"],"color_identity":["W"],"keywords":["Aftermath"],"card_faces":[{"object":"card_face","name":"Dusk","mana_cost":"{2}{W}{W}","type_line":"Sorcery","oracle_text":"Destroy all creatures with power 3 or greater.","artist":"Noah Bradley","artist_id":"81995d11-da98-4f8b-89bd-b88ca2ddb06b","illustration_id":"f3d63aed-2784-4ef5-9676-846b1e65e040"},{"object":"card_face","name":"Dawn","mana_cost":"{3}{W}{W}","type_line":"Sorcery","oracle_text":"Aftermath (Cast this spell only from your graveyard. Then exile it.)\nReturn all creature cards with power 2 or less from your graveyard to your hand.","artist":"Noah Bradley","artist_id":"81995d11-da98-4f8b-89bd-b88ca2ddb06b"}],"legalities":{"standard":"not_legal","future":"not_legal","historic":"legal","timeless":"legal","gladiator":"legal","pioneer":"legal","explorer":"legal","modern":"legal","legacy":"legal","pauper":"not_legal","vintage":"legal","penny":"legal","commander":"legal","oathbreaker":"legal","standardbrawl":"not_legal","brawl":"legal","alchemy":"not_legal","paupercommander":"not_legal","duel":"legal","oldschool":"not_legal","premodern":"not_legal","predh":"not_legal"},"games":["paper","mtgo"],"reserved":false,"game_changer":false,"foil":true,"nonfoil":true,"finishes":["nonfoil","foil"],"oversized":false,"promo":false,"reprint":false,"variation":false,"set_id":"02d1c536-68bc-4208-9b65-7741ef1f9da8","set":"akh","set_name":"Amonkhet","set_type":"expansion","set_uri":"https://api.scryfall.com/sets/02d1c536-68bc-4208-9b65-7741ef1f9da8","set_search_uri":"https://api.scryfall.com/cards/search?order=set&q=e%3Aakh&unique=prints","scryfall_set_uri":"https://scryfall.com/sets/akh?utm_source=api","rulings_uri":"https://api.scryfall.com/cards/937dbc51-b589-4237-9fce-ea5c757f7c48/rulings","prints_search_uri":"https://api.scryfall.com/cards/search?order=released&q=oracleid%3A7bc3f92f-68a2-4934-afc4-89f6d0e8cf98&unique=prints","collector_number":"210","digital":false,"rarity":"rare","card_back_id":"0aeebaf5-8c7d-4636-9e82-8c27447861f7","artist":"Noah Bradley","artist_ids":["81995d11-da98-4f8b-89bd-b88ca2ddb06b"],"illustration_id":"f3d63aed-2784-4ef5-9676-846b1e65e040","border_color":"black","frame":"2015","security_stamp":"oval","full_art":false,"textless":false,"booster":true,"story_spotlight":false,"edhrec_rank":830,"penny_rank":3788,"prices":{"usd":"0.35","usd_foil":"4.17","usd_etched":null,"eur":"0.54","eur_foil":"1.55","tix":"0.02"},"related_uris":{"gatherer":"https://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=426912&printed=false","tcgplayer_infinite_articles":"https://partner.tcgplayer.com/c/4931599/1830156/21018?subId1=api&trafcat=infinite&u=https%3A%2F%2Finfinite.tcgplayer.com%2Fsearch%3FcontentMode%3Darticle%26game%3Dmagic%26q%3DDusk%2B%252F%252F%2BDawn","tcgplayer_infinite_decks":"https://partner.tcgplayer.com/c/4931599/1830156/21018?subId1=api&trafcat=infinite&u=https%3A%2F%2Finfinite.tcgplayer.com%2Fsearch%3FcontentMode%3Ddeck%26game%3Dmagic%26q%3DDusk%2B%252F%252F%2BDawn","edhrec":"https://edhrec.com/route/?cc=Dusk+%2F%2F+Dawn"},"purchase_uris":{"tcgplayer":"https://partner.tcgplayer.com/c/4931599/1830156/21018?subId1=api&u=https%3A%2F%2Fwww.tcgplayer.com%2Fproduct%2F129823%3Fpage%3D1","cardmarket":"https://www.cardmarket.com/en/Magic/Products/Singles/Amonkhet/Dusk-Dawn?referrer=scryfall&utm_campaign=card_prices&utm_medium=text&utm_source=scryfall","cardhoarder":"https://www.cardhoarder.com/cards/64026?affiliate_id=scryfall&ref=card-profile&utm_campaign=affiliate&utm_medium=card&utm_source=scryfall"}}`

var duskDawn = Card{
	ID:            "937dbc51-b589-4237-9fce-ea5c757f7c48",
	OracleID:      "7bc3f92f-68a2-4934-afc4-89f6d0e8cf98",
	MultiverseIDs: []int{426912},
	MTGOID:        intPointer(64026),
	MTGOFoilID:    nil,
	Name:          "Dusk // Dawn",
	Lang:          LangEnglish,
	URI:           "https://api.scryfall.com/cards/937dbc51-b589-4237-9fce-ea5c757f7c48",
	ScryfallURI:   "https://scryfall.com/card/akh/210/dusk-dawn?utm_source=api",
	TCGPlayerID:   intPointer(129823),
	Layout:        LayoutSplit,
	HighresImage:  true,
	ImageURIs: &ImageURIs{
		Small:      "https://cards.scryfall.io/small/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330",
		Normal:     "https://cards.scryfall.io/normal/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330",
		Large:      "https://cards.scryfall.io/large/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330",
		PNG:        "https://cards.scryfall.io/png/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.png?1549941330",
		ArtCrop:    "https://cards.scryfall.io/art_crop/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330",
		BorderCrop: "https://cards.scryfall.io/border_crop/front/9/3/937dbc51-b589-4237-9fce-ea5c757f7c48.jpg?1549941330",
	},
	CMC:           9,
	TypeLine:      "Sorcery // Sorcery",
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
		Standard:  "not_legal",
		Modern:    "legal",
		Pauper:    "not_legal",
		Pioneer:   "legal",
		Legacy:    "legal",
		Penny:     "legal",
		Vintage:   "legal",
		Duel:      "legal",
		Commander: "legal",
		Future:    "not_legal",
	},
	Reserved:        false,
	Foil:            true,
	NonFoil:         true,
	Reprint:         false,
	Set:             "akh",
	SetName:         "Amonkhet",
	SetURI:          "https://api.scryfall.com/sets/02d1c536-68bc-4208-9b65-7741ef1f9da8",
	SetSearchURI:    "https://api.scryfall.com/cards/search?order=set&q=e%3Aakh&unique=prints",
	ScryfallSetURI:  "https://scryfall.com/sets/akh?utm_source=api",
	RulingsURI:      "https://api.scryfall.com/cards/937dbc51-b589-4237-9fce-ea5c757f7c48/rulings",
	PrintsSearchURI: "https://api.scryfall.com/cards/search?order=released&q=oracleid%3A7bc3f92f-68a2-4934-afc4-89f6d0e8cf98&unique=prints",
	CollectorNumber: "210",
	Digital:         false,
	Rarity:          "rare",
	IllustrationID:  stringPointer("f3d63aed-2784-4ef5-9676-846b1e65e040"),
	Artist:          stringPointer("Noah Bradley"),
	Frame:           Frame2015,
	FrameEffects:    nil,
	FullArt:         false,
	BorderColor:     "black",
	EDHRECRank:      intPointer(830),
	Prices: Prices{
		USD:     "0.35",
		USDFoil: "4.17",
		EUR:     "0.54",
		EURFoil: "1.55",
		Tix:     "0.02",
	},
	RelatedURIs: RelatedURIs{
		Gatherer:       "https://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=426912&printed=false",
		TCGPlayerDecks: "",
		EDHREC:         "https://edhrec.com/route/?cc=Dusk+%2F%2F+Dawn",
		MTGTop8:        "",
	},
	ReleasedAt: Date{Time: time.Date(2017, 04, 28, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
	PurchaseURIs: PurchaseURIs{
		TCGPlayer:   "https://partner.tcgplayer.com/c/4931599/1830156/21018?subId1=api&u=https%3A%2F%2Fwww.tcgplayer.com%2Fproduct%2F129823%3Fpage%3D1",
		CardMarket:  "https://www.cardmarket.com/en/Magic/Products/Singles/Amonkhet/Dusk-Dawn?referrer=scryfall&utm_campaign=card_prices&utm_medium=text&utm_source=scryfall",
		CardHoarder: "https://www.cardhoarder.com/cards/64026?affiliate_id=scryfall&ref=card-profile&utm_campaign=affiliate&utm_medium=card&utm_source=scryfall",
	},
	Keywords:    []string{"Aftermath"},
	Booster:     true,
	Finishes:    []Finish{FinishNonFoil, FinishFoil},
	ImageStatus: (*ImageStatus)(stringPointer(string(ImageStatusHighres))),
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
		w.Write([]byte(`{"object": "list", "total_cards": 1000, "has_more": true, "next_page": "https://api.scryfall.com/cards?page=3", "data": [` + duskDawnJSON + `]}`))
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
		w.Write([]byte(duskDawnJSON))
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

		w.Write([]byte(`{"object": "catalog", "total_items": 20, "data": ["Thallid", "Thorn Thallid", "Thalakos Seer", "Thalakos Scout", "Thalia's Lancers", "Thalakos Sentry", "Thallid Devourer", "Thalakos Deceiver", "Thalakos Drifters", "Thalakos Lowlands", "Thalakos Mistfolk", "Thallid Soothsayer", "Thallid Germinator", "Thalia's Lieutenant", "Thallid Shell-Dweller", "Thalia, Heretic Cathar", "Thalakos Dreamsower", "Thalia, Guardian of Thraben", "Tukatongue Thallid", "Lethal Sting"]}`))
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
		w.Write([]byte(duskDawnJSON))
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
		w.Write([]byte(duskDawnJSON))
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
		w.Write([]byte(duskDawnJSON))
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
		w.Write([]byte(duskDawnJSON))
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
		w.Write([]byte(duskDawnJSON))
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
		w.Write([]byte(duskDawnJSON))
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
