package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCardNamesCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/card-names", "total_values": 7, "data": ["\"Ach! Hans, Run!\"", "\"Rumors of My Death . . .\"", "1996 World Champion", "A Display of My Dark Power", "A Reckoning Approaches", "AWOL", "Abandon Hope"]}`)
	})
	client, ts, err := setupTestServer("/catalog/card-names", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetCardNamesCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/card-names",
		TotalValues: 7,
		Data: []string{
			"\"Ach! Hans, Run!\"",
			"\"Rumors of My Death . . .\"",
			"1996 World Champion",
			"A Display of My Dark Power",
			"A Reckoning Approaches",
			"AWOL",
			"Abandon Hope",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}

func TestGetWordBankCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/word-bank", "total_values": 11, "data": ["abandon", "abandoned", "abattoir", "abbey", "abbot", "abc's", "abdallah", "abduction", "aberrant", "aberration", "abeyance"]}`)
	})
	client, ts, err := setupTestServer("/catalog/word-bank", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetWordBankCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/word-bank",
		TotalValues: 11,
		Data: []string{
			"abandon",
			"abandoned",
			"abattoir",
			"abbey",
			"abbot",
			"abc's",
			"abdallah",
			"abduction",
			"aberrant",
			"aberration",
			"abeyance",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}

func TestGetCreatureTypesCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/creature-types", "total_values": 11, "data": ["Advisor", "Aetherborn", "Ally", "Angel", "Antelope", "Ape", "Archer", "Archon", "Artificer", "Assassin", "Assembly-Worker", "Atog"]}`)
	})
	client, ts, err := setupTestServer("/catalog/creature-types", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetCreatureTypesCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/creature-types",
		TotalValues: 11,
		Data: []string{
			"Advisor",
			"Aetherborn",
			"Ally",
			"Angel",
			"Antelope",
			"Ape",
			"Archer",
			"Archon",
			"Artificer",
			"Assassin",
			"Assembly-Worker",
			"Atog",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}

func TestGetPlaneswalkerTypesCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/planeswalker-types", "total_values": 10, "data": ["Ajani", "Angrath", "Arlinn", "Ashiok", "Bolas", "Chandra", "Dack", "Daretti", "Domri", "Dovin"]}`)
	})
	client, ts, err := setupTestServer("/catalog/planeswalker-types", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetPlaneswalkerTypesCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/planeswalker-types",
		TotalValues: 10,
		Data: []string{
			"Ajani",
			"Angrath",
			"Arlinn",
			"Ashiok",
			"Bolas",
			"Chandra",
			"Dack",
			"Daretti",
			"Domri",
			"Dovin",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}

func TestGetLandTypesCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/land-types", "total_values": 13, "data": ["Desert", "Forest", "Gate", "Island", "Lair", "Locus", "Mine", "Mountain", "Plains", "Power-Plant", "Swamp", "Tower", "Urza’s"]}`)
	})
	client, ts, err := setupTestServer("/catalog/land-types", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetLandTypesCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/land-types",
		TotalValues: 13,
		Data: []string{
			"Desert",
			"Forest",
			"Gate",
			"Island",
			"Lair",
			"Locus",
			"Mine",
			"Mountain",
			"Plains",
			"Power-Plant",
			"Swamp",
			"Tower",
			"Urza’s",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}

func TestGetSpellTypesCatalog(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/spell-types", "total_values": 2, "data": ["Arcane", "Trap"]}`)
	})
	client, ts, err := setupTestServer("/catalog/spell-types", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	catalog, err := client.GetSpellTypesCatalog(ctx)
	if err != nil {
		t.Fatalf("Error getting catalog: %v", err)
	}

	want := Catalog{
		URI:         "https://api.scryfall.com/catalog/spell-types",
		TotalValues: 2,
		Data: []string{
			"Arcane",
			"Trap",
		},
	}
	if !reflect.DeepEqual(catalog, want) {
		t.Errorf("got: %#v want: %#v", catalog, want)
	}
}
