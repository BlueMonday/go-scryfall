package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetCardNamesCatalog(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/catalog/card-names", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/card-names", "total_values": 7, "data": ["\"Ach! Hans, Run!\"", "\"Rumors of My Death . . .\"", "1996 World Champion", "A Display of My Dark Power", "A Reckoning Approaches", "AWOL", "Abandon Hope"]}`)
	}))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	client, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("Error creating new client: %v", err)
	}

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
	mux := http.NewServeMux()
	mux.HandleFunc("/catalog/word-bank", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "catalog", "uri": "https://api.scryfall.com/catalog/word-bank", "total_values": 11, "data": ["abandon", "abandoned", "abattoir", "abbey", "abbot", "abc's", "abdallah", "abduction", "aberrant", "aberration", "abeyance"]}`)
	}))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	client, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("Error creating new client: %v", err)
	}

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
