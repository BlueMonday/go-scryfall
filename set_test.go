package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestListSets(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "set", "code": "dom", "mtgo_code": "dom", "name": "Dominaria", "uri": "https://api.scryfall.com/sets/dom", "scryfall_uri": "https://scryfall.com/sets/dom", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Adom&unique=prints", "released_at": "2018-04-27", "set_type": "expansion", "card_count": 142, "digital": false, "foil": false, "icon_svg_uri": "https://assets.scryfall.com/assets/sets/dom.svg"}, {"object": "set", "code": "a25", "name": "Masters 25", "uri": "https://api.scryfall.com/sets/a25", "scryfall_uri": "https://scryfall.com/sets/a25", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Aa25&unique=prints", "released_at": "2018-03-16", "set_type": "masters", "card_count": 249, "digital": false, "foil": false, "icon_svg_uri": "https://assets.scryfall.com/assets/sets/a25.svg"},{"object":"set","code":"rix","mtgo_code":"rix","name":"Rivals of Ixalan","uri":"https://api.scryfall.com/sets/rix","scryfall_uri":"https://scryfall.com/sets/rix","search_uri":"https://api.scryfall.com/cards/search?order=set\u0026q=e%3Arix\u0026unique=prints","released_at":"2018-01-19","set_type":"expansion","card_count":205,"digital":false,"foil":false,"block_code":"xln","block":"Ixalan","icon_svg_uri":"https://assets.scryfall.com/assets/sets/rix.svg"}]}`)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("Error creating new client: %v", err)
	}

	ctx := context.Background()
	sets, err := client.ListSets(ctx)
	if err != nil {
		t.Fatalf("Error listing sets: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT-8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	ixalanBlockCode := "xln"
	ixalanBlock := "Ixalan"
	want := []Set{
		{
			Code:        "dom",
			MTGOCode:    "dom",
			Name:        "Dominaria",
			URI:         "https://api.scryfall.com/sets/dom",
			ScryfallURI: "https://scryfall.com/sets/dom",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Adom&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 04, 27, 0, 0, 0, 0, loc)},
			SetType:     "expansion",
			CardCount:   142,
			Digital:     false,
			Foil:        false,
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/dom.svg",
		},
		{
			Code:        "a25",
			MTGOCode:    "",
			Name:        "Masters 25",
			URI:         "https://api.scryfall.com/sets/a25",
			ScryfallURI: "https://scryfall.com/sets/a25",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Aa25&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 03, 16, 0, 0, 0, 0, loc)},
			SetType:     "masters",
			CardCount:   249,
			Digital:     false,
			Foil:        false,
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/a25.svg",
		},
		{
			Code:        "rix",
			MTGOCode:    "rix",
			Name:        "Rivals of Ixalan",
			URI:         "https://api.scryfall.com/sets/rix",
			ScryfallURI: "https://scryfall.com/sets/rix",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Arix&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 01, 19, 0, 0, 0, 0, loc)},
			SetType:     "expansion",
			CardCount:   205,
			Digital:     false,
			Foil:        false,
			BlockCode:   &ixalanBlockCode,
			Block:       &ixalanBlock,
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/rix.svg",
		},
	}

	if !reflect.DeepEqual(sets, want) {
		t.Errorf("got: %#v want: %#v", sets, want)
	}
}

func TestGetSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "set", "code": "aer", "mtgo_code": "aer", "name": "Aether Revolt", "uri": "https://api.scryfall.com/sets/aer", "scryfall_uri": "https://scryfall.com/sets/aer", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Aaer&unique=prints", "released_at": "2017-01-20", "set_type": "expansion", "card_count": 194, "digital": false, "foil": false, "block_code": "kld", "block": "Kaladesh", "icon_svg_uri": "https://assets.scryfall.com/assets/sets/aer.svg"}`)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("Error creating new client: %v", err)
	}

	ctx := context.Background()
	set, err := client.GetSet(ctx, "aer")
	if err != nil {
		t.Fatalf("Error getting set: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT-8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	aetherRevoltBlockCode := "kld"
	aetherRevoltBlock := "Kaladesh"
	want := Set{
		Code:        "aer",
		MTGOCode:    "aer",
		Name:        "Aether Revolt",
		URI:         "https://api.scryfall.com/sets/aer",
		ScryfallURI: "https://scryfall.com/sets/aer",
		SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Aaer&unique=prints",
		ReleasedAt:  &Date{Time: time.Date(2017, 01, 20, 0, 0, 0, 0, loc)},
		SetType:     "expansion",
		CardCount:   194,
		Digital:     false,
		Foil:        false,
		BlockCode:   &aetherRevoltBlockCode,
		Block:       &aetherRevoltBlock,
		IconSVGURI:  "https://assets.scryfall.com/assets/sets/aer.svg",
	}

	if !reflect.DeepEqual(set, want) {
		t.Errorf("got: %#v want: %#v", set, want)
	}
}
