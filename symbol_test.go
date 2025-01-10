package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListCardSymbols(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "card_symbol", "symbol": "{T}", "svg_uri": "https://svgs.scryfall.io/card-symbols/T.svg", "loose_variant": null, "english": "tap this permanent", "transposable": false, "represents_mana": false, "appears_in_mana_costs": false, "mana_value": 0, "hybrid": false, "phyrexian": false, "cmc": 0, "funny": false, "colors": [], "gatherer_alternates": ["ocT", "oT"]}, {"object": "card_symbol", "symbol": "{Q}", "svg_uri": "https://svgs.scryfall.io/card-symbols/Q.svg", "loose_variant": null, "english": "untap this permanent", "transposable": false, "represents_mana": false, "appears_in_mana_costs": false, "mana_value": 0, "hybrid": false, "phyrexian": false, "cmc": 0, "funny": false, "colors": [], "gatherer_alternates": null}, {"object": "card_symbol", "symbol": "{W}", "svg_uri": "https://svgs.scryfall.io/card-symbols/W.svg", "loose_variant": "W", "english": "one white mana", "transposable": false, "represents_mana": true, "appears_in_mana_costs": true, "mana_value": 1, "hybrid": false, "phyrexian": false, "cmc": 1, "funny": false, "colors": ["W"], "gatherer_alternates": ["oW", "ooW"]}]}`)
	})
	client, ts, err := setupTestServer("/symbology", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	cardSymbols, err := client.ListCardSymbols(ctx)
	if err != nil {
		t.Fatalf("Error listing card symbols: %v", err)
	}

	want := []CardSymbol{
		{
			Object:             "card_symbol",
			Symbol:             "{T}",
			SVGURI:             stringPointer("https://svgs.scryfall.io/card-symbols/T.svg"),
			LooseVariant:       nil,
			English:            "tap this permanent",
			Transposable:       false,
			RepresentsMana:     false,
			AppearsInManaCosts: false,
			ManaValue:          float64Pointer(0),
			Hybrid:             false,
			Phyrexian:          false,
			CMC:                0,
			Funny:              false,
			Colors:             []Color{},
			GathererAlternates: []string{"ocT", "oT"},
		},
		{
			Object:             "card_symbol",
			Symbol:             "{Q}",
			SVGURI:             stringPointer("https://svgs.scryfall.io/card-symbols/Q.svg"),
			LooseVariant:       nil,
			English:            "untap this permanent",
			Transposable:       false,
			AppearsInManaCosts: false,
			RepresentsMana:     false,
			ManaValue:          float64Pointer(0),
			Hybrid:             false,
			Phyrexian:          false,
			CMC:                0,
			Funny:              false,
			Colors:             []Color{},
			GathererAlternates: nil,
		},
		{
			Object:             "card_symbol",
			Symbol:             "{W}",
			SVGURI:             stringPointer("https://svgs.scryfall.io/card-symbols/W.svg"),
			LooseVariant:       stringPointer("W"),
			English:            "one white mana",
			Transposable:       false,
			AppearsInManaCosts: true,
			RepresentsMana:     true,
			ManaValue:          float64Pointer(1),
			Hybrid:             false,
			Phyrexian:          false,
			CMC:                1,
			Funny:              false,
			Colors:             []Color{ColorWhite},
			GathererAlternates: []string{"oW", "ooW"},
		},
	}
	if !reflect.DeepEqual(cardSymbols, want) {
		t.Errorf("got: %#v want: %#v", cardSymbols, want)
	}
}

func TestParseManaCost(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cost := r.URL.Query().Get("cost")
		if cost != "RUx" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, `{"object": "mana_cost", "cost": "{X}{U}{R}", "colors": ["U", "R"], "cmc": 2, "colorless": false, "monocolored": false, "multicolored": true}`)
	})
	client, ts, err := setupTestServer("/symbology/parse-mana", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	manaCost, err := client.ParseManaCost(ctx, "RUx")
	if err != nil {
		t.Fatalf("Error parsing mana cost: %v", err)
	}

	want := ManaCost{
		Cost:         "{X}{U}{R}",
		CMC:          2,
		Colorless:    false,
		Monocolored:  false,
		Multicolored: true,
		Colors:       []Color{ColorBlue, ColorRed},
	}
	if !reflect.DeepEqual(manaCost, want) {
		t.Errorf("got: %#v want: %#v", manaCost, want)
	}
}
