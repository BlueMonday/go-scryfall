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
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "card_symbol", "symbol": "{T}", "loose_variant": null, "english": "tap this permanent", "transposable": false, "represents_mana": false, "appears_in_mana_costs": false, "cmc": 0, "funny": false, "colors": []}, {"object": "card_symbol", "symbol": "{Q}", "loose_variant": null, "english": "untap this permanent", "transposable": false, "represents_mana": false, "appears_in_mana_costs": false, "cmc": 0, "funny": false, "colors": []}, {"object": "card_symbol", "symbol": "{W/U}", "loose_variant": null, "english": "one white or blue mana", "transposable": true, "represents_mana": true, "appears_in_mana_costs": true, "cmc": 1, "funny": false, "colors": ["W", "U"]}]}`)
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
			Symbol:             "{T}",
			LooseVariant:       nil,
			English:            "tap this permanent",
			Transposable:       false,
			RepresentsMana:     false,
			CMC:                0,
			AppearsInManaCosts: false,
			Funny:              false,
			Colors:             []Color{},
		},
		{
			Symbol:             "{Q}",
			LooseVariant:       nil,
			English:            "untap this permanent",
			Transposable:       false,
			RepresentsMana:     false,
			CMC:                0,
			AppearsInManaCosts: false,
			Funny:              false,
			Colors:             []Color{},
		},
		{
			Symbol:             "{W/U}",
			LooseVariant:       nil,
			English:            "one white or blue mana",
			Transposable:       true,
			RepresentsMana:     true,
			CMC:                1,
			AppearsInManaCosts: true,
			Funny:              false,
			Colors:             []Color{ColorWhite, ColorBlue},
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
