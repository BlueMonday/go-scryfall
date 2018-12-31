package scryfall

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestGetRulingsByMultiverseID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "ruling", "source": "wotc", "published_at": "2004-10-04", "comment": "The ability is a mana ability, so it is activated and resolves as a mana ability, but it can only be activated at times when you can cast an instant. Yes, this is a bit weird."}]}`)
	})
	client, ts, err := setupTestServer("/cards/multiverse/3255/rulings", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	rulings, err := client.GetRulingsByMultiverseID(ctx, 3255)
	if err != nil {
		t.Fatalf("Error getting rulings: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT+8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	want := []Ruling{
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2004, 10, 04, 0, 0, 0, 0, loc)},
			Comment:     "The ability is a mana ability, so it is activated and resolves as a mana ability, but it can only be activated at times when you can cast an instant. Yes, this is a bit weird.",
		},
	}
	if !reflect.DeepEqual(rulings, want) {
		t.Errorf("got: %#v want: %#v", rulings, want)
	}
}

func TestGetRulingsByMTGOID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "You choose the mode as the triggered ability goes on the stack. You can choose a mode that requires targets only if there are legal targets available."}, {"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "If the ability is countered (either for having its target become illegal or because a spell or ability counters it), the mode chosen for that instance of the ability still counts as being chosen."}, {"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "The phrase “that hasn’t been chosen” refers only to that specific Demonic Pact. If you control one and cast another one, you can choose any mode for the second one the first time its ability triggers."}, {"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "It doesn’t matter who has chosen any particular mode. For example, say you control Demonic Pact and have chosen the first two modes. If an opponent gains control of Demonic Pact, that player can choose only the third or fourth mode."}, {"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "In some very unusual situations, you may not be able to choose a mode, either because all modes have previously been chosen or the only remaining modes require targets and there are no legal targets available. In this case, the ability is simply removed from the stack with no effect."}, {"object": "ruling", "source": "wotc", "published_at": "2015-06-22", "comment": "Yes, if the fourth mode is the only one remaining, you must choose it. You read the whole contract, right?"}]}`)
	})
	client, ts, err := setupTestServer("/cards/mtgo/57934/rulings", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	rulings, err := client.GetRulingsByMTGOID(ctx, 57934)
	if err != nil {
		t.Fatalf("Error getting rulings: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT+8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	want := []Ruling{
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "You choose the mode as the triggered ability goes on the stack. You can choose a mode that requires targets only if there are legal targets available.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "If the ability is countered (either for having its target become illegal or because a spell or ability counters it), the mode chosen for that instance of the ability still counts as being chosen.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "The phrase “that hasn’t been chosen” refers only to that specific Demonic Pact. If you control one and cast another one, you can choose any mode for the second one the first time its ability triggers.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "It doesn’t matter who has chosen any particular mode. For example, say you control Demonic Pact and have chosen the first two modes. If an opponent gains control of Demonic Pact, that player can choose only the third or fourth mode.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "In some very unusual situations, you may not be able to choose a mode, either because all modes have previously been chosen or the only remaining modes require targets and there are no legal targets available. In this case, the ability is simply removed from the stack with no effect.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2015, 06, 22, 0, 0, 0, 0, loc)},
			Comment:     "Yes, if the fourth mode is the only one remaining, you must choose it. You read the whole contract, right?",
		},
	}
	if !reflect.DeepEqual(rulings, want) {
		t.Errorf("got: %#v want: %#v", rulings, want)
	}
}

func TestGetRulingsBySetCodeAndCollectorNumber(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "ruling", "source": "wotc", "published_at": "2017-11-17", "comment": "Mana Drain can target a spell that can’t be countered. When Mana Drain resolves, that spell won’t be countered, but you’ll still add mana to your mana pool at the beginning of your next main phase."}, {"object": "ruling", "source": "wotc", "published_at": "2017-11-17", "comment": "If the target spell is an illegal target when Mana Drain tries to resolve, it will be countered and none of its effects will happen. You won’t get any mana."}, {"object": "ruling", "source": "wotc", "published_at": "2017-11-17", "comment": "Mana Drain’s delayed triggered ability will usually trigger at the beginning of your precombat main phase. However, if you cast Mana Drain during your precombat main phase or during your combat phase, its delayed triggered ability will trigger at the beginning of that turn’s postcombat main phase."}]}`)
	})
	client, ts, err := setupTestServer("/cards/ima/65/rulings", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	rulings, err := client.GetRulingsBySetCodeAndCollectorNumber(ctx, "ima", 65)
	if err != nil {
		t.Fatalf("Error getting rulings: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT+8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	want := []Ruling{
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2017, 11, 17, 0, 0, 0, 0, loc)},
			Comment:     "Mana Drain can target a spell that can’t be countered. When Mana Drain resolves, that spell won’t be countered, but you’ll still add mana to your mana pool at the beginning of your next main phase.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2017, 11, 17, 0, 0, 0, 0, loc)},
			Comment:     "If the target spell is an illegal target when Mana Drain tries to resolve, it will be countered and none of its effects will happen. You won’t get any mana.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2017, 11, 17, 0, 0, 0, 0, loc)},
			Comment:     "Mana Drain’s delayed triggered ability will usually trigger at the beginning of your precombat main phase. However, if you cast Mana Drain during your precombat main phase or during your combat phase, its delayed triggered ability will trigger at the beginning of that turn’s postcombat main phase.",
		},
	}
	if !reflect.DeepEqual(rulings, want) {
		t.Errorf("got: %#v want: %#v", rulings, want)
	}
}

func TestGetRulings(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": [{"object": "ruling", "source": "wotc", "published_at": "2004-10-04", "comment": "It must flip like a coin and not like a Frisbee."}, {"object": "ruling", "source": "wotc", "published_at": "2004-10-04", "comment": "Only cards touched when it stops moving are affected. Not ones touched while it is moving."}]}`)
	})
	client, ts, err := setupTestServer("/cards/f2b9983e-20d4-4d12-9e2c-ec6d9a345787/rulings", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	rulings, err := client.GetRulings(ctx, "f2b9983e-20d4-4d12-9e2c-ec6d9a345787")
	if err != nil {
		t.Fatalf("Error getting rulings: %v", err)
	}

	loc, err := time.LoadLocation("Etc/GMT+8")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}

	want := []Ruling{
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2004, 10, 4, 0, 0, 0, 0, loc)},
			Comment:     "It must flip like a coin and not like a Frisbee.",
		},
		{
			Source:      SourceWOTC,
			PublishedAt: Date{Time: time.Date(2004, 10, 4, 0, 0, 0, 0, loc)},
			Comment:     "Only cards touched when it stops moving are affected. Not ones touched while it is moving.",
		},
	}
	if !reflect.DeepEqual(rulings, want) {
		t.Errorf("got: %#v want: %#v", rulings, want)
	}
}
