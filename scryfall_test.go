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

func stringPointer(v string) *string {
	return &v
}

func intPointer(v int) *int {
	return &v
}

func TestDateUnmarshalJSON(t *testing.T) {
	loc, err := time.LoadLocation("Etc/GMT-8")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		in  []byte
		out Date
	}{
		{
			[]byte("null"),
			Date{Time: time.Time{}},
		},
		{
			[]byte("2018-04-27"),
			Date{Time: time.Date(2018, 04, 27, 0, 0, 0, 0, loc)},
		},
	}

	for _, test := range tests {
		t.Run(string(test.in), func(t *testing.T) {
			date := Date{}
			err := date.UnmarshalJSON(test.in)
			if err != nil {
				t.Fatalf("Unexpected error while unmarshaling JSON date representation: %v", err)
			}

			if !reflect.DeepEqual(date, test.out) {
				t.Errorf("got: %s want: %s", date, test.out)
			}
		})
	}
}

func TestError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/cards/nope", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"object": "error", "code": "not_found", "status": 404, "details": "The requested object or REST method was not found."}`)
	}))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	client, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("Error creating new client: %v", err)
	}

	ctx := context.Background()
	_, err = client.GetCard(ctx, "nope")

	expectedErr := &Error{
		Code:    "not_found",
		Status:  404,
		Details: "The requested object or REST method was not found.",
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("got: %#v want: %#v", err, expectedErr)
	}
}
