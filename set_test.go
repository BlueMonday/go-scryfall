package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListSets(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "set", "code": "dom", "mtgo_code": "dar", "arena_code": "dar", "name": "Dominaria", "uri": "https://api.scryfall.com/sets/dom", "scryfall_uri": "https://scryfall.com/sets/dom", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Adom&unique=prints", "released_at": "2018-04-27", "set_type": "expansion", "card_count": 142, "digital": false, "foil": false, "icon_svg_uri": "https://assets.scryfall.com/assets/sets/dom.svg"}, {"object": "set", "code": "a25", "mtgo_code": "a25", "arena_code": "a25", "name": "Masters 25", "uri": "https://api.scryfall.com/sets/a25", "scryfall_uri": "https://scryfall.com/sets/a25", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Aa25&unique=prints", "released_at": "2018-03-16", "set_type": "masters", "card_count": 249, "digital": false, "foil": false, "icon_svg_uri": "https://assets.scryfall.com/assets/sets/a25.svg"},{"object":"set","code":"rix","mtgo_code":"rix","arena_code":"rix","name":"Rivals of Ixalan","uri":"https://api.scryfall.com/sets/rix","scryfall_uri":"https://scryfall.com/sets/rix","search_uri":"https://api.scryfall.com/cards/search?order=set\u0026q=e%3Arix\u0026unique=prints","released_at":"2018-01-19","set_type":"expansion","card_count":205,"digital":false,"foil":false,"block_code":"xln","block":"Ixalan","icon_svg_uri":"https://assets.scryfall.com/assets/sets/rix.svg"}]}`)
	})
	client, ts, err := setupTestServer("/sets", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	sets, err := client.ListSets(ctx)
	if err != nil {
		t.Fatalf("Error listing sets: %v", err)
	}

	mtgoCodes := []string{
		"dar",
		"a25",
		"rix",
	}
	arenaCodes := []string{
		"dar",
		"a25",
		"rix",
	}
	want := []Set{
		{
			Code:        "dom",
			MTGOCode:    &mtgoCodes[0],
			ArenaCode:   &arenaCodes[0],
			Name:        "Dominaria",
			URI:         "https://api.scryfall.com/sets/dom",
			ScryfallURI: "https://scryfall.com/sets/dom",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Adom&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 04, 27, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
			SetType:     "expansion",
			CardCount:   142,
			Digital:     false,
			FoilOnly:    false,
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/dom.svg",
		},
		{
			Code:        "a25",
			MTGOCode:    &mtgoCodes[1],
			ArenaCode:   &arenaCodes[1],
			Name:        "Masters 25",
			URI:         "https://api.scryfall.com/sets/a25",
			ScryfallURI: "https://scryfall.com/sets/a25",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Aa25&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 03, 16, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
			SetType:     "masters",
			CardCount:   249,
			Digital:     false,
			FoilOnly:    false,
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/a25.svg",
		},
		{
			Code:        "rix",
			MTGOCode:    &mtgoCodes[2],
			ArenaCode:   &arenaCodes[2],
			Name:        "Rivals of Ixalan",
			URI:         "https://api.scryfall.com/sets/rix",
			ScryfallURI: "https://scryfall.com/sets/rix",
			SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Arix&unique=prints",
			ReleasedAt:  &Date{Time: time.Date(2018, 01, 19, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
			SetType:     "expansion",
			CardCount:   205,
			Digital:     false,
			FoilOnly:    false,
			BlockCode:   stringPointer("xln"),
			Block:       stringPointer("Ixalan"),
			IconSVGURI:  "https://assets.scryfall.com/assets/sets/rix.svg",
		},
	}
	if !reflect.DeepEqual(sets, want) {
		t.Errorf("got: %#v want: %#v", sets, want)
	}
}

func TestGetSet(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "set", "code": "aer", "mtgo_code": "aer", "name": "Aether Revolt", "uri": "https://api.scryfall.com/sets/aer", "scryfall_uri": "https://scryfall.com/sets/aer", "search_uri": "https://api.scryfall.com/cards/search?order=set&q=e%3Aaer&unique=prints", "released_at": "2017-01-20", "set_type": "expansion", "card_count": 194, "digital": false, "foil": false, "block_code": "kld", "block": "Kaladesh", "icon_svg_uri": "https://assets.scryfall.com/assets/sets/aer.svg"}`)
	})
	client, ts, err := setupTestServer("/sets/aer", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	set, err := client.GetSet(ctx, "aer")
	if err != nil {
		t.Fatalf("Error getting set: %v", err)
	}

	aetherSetCode := "aer"
	aetherRevoltBlockCode := "kld"
	aetherRevoltBlock := "Kaladesh"
	want := Set{
		Code:        aetherSetCode,
		MTGOCode:    &aetherSetCode,
		ArenaCode:   nil,
		Name:        "Aether Revolt",
		URI:         "https://api.scryfall.com/sets/aer",
		ScryfallURI: "https://scryfall.com/sets/aer",
		SearchURI:   "https://api.scryfall.com/cards/search?order=set&q=e%3Aaer&unique=prints",
		ReleasedAt:  &Date{Time: time.Date(2017, 01, 20, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
		SetType:     "expansion",
		CardCount:   194,
		Digital:     false,
		FoilOnly:    false,
		BlockCode:   &aetherRevoltBlockCode,
		Block:       &aetherRevoltBlock,
		IconSVGURI:  "https://assets.scryfall.com/assets/sets/aer.svg",
	}
	if !reflect.DeepEqual(set, want) {
		t.Errorf("got: %#v want: %#v", set, want)
	}
}
