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

func setupTestServer(pattern string, handler func(http.ResponseWriter, *http.Request), clientOptions ...ClientOption) (*Client, *httptest.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, handler)
	ts := httptest.NewServer(mux)

	mergedClientOptions := []ClientOption{WithBaseURL(ts.URL), WithLimiter(nil)}
	mergedClientOptions = append(mergedClientOptions, clientOptions...)
	client, err := NewClient(mergedClientOptions...)
	if err != nil {
		ts.Close()
		return nil, nil, err
	}

	return client, ts, nil
}

func TestDateUnmarshalJSON(t *testing.T) {

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
			Date{Time: time.Date(2018, 4, 27, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))},
		},
	}

	for _, test := range tests {
		t.Run(string(test.in), func(t *testing.T) {
			date := Date{}
			err := date.UnmarshalJSON(test.in)
			if err != nil {
				t.Fatalf("Unexpected error while unmarshaling JSON date representation: %v", err)
			}

			if !date.Time.Equal(test.out.Time) {
				t.Errorf("got: %s want: %s", date, test.out)
			}
		})
	}
}

func TestTimestampUnmarshalJSON(t *testing.T) {
	tests := []struct {
		in  []byte
		out Timestamp
	}{
		{
			[]byte("null"),
			Timestamp{Time: time.Time{}},
		},
		{
			[]byte("2018-12-01T14:31:43-05:00"),
			Timestamp{Time: time.Date(2018, 12, 1, 14, 31, 43, 0, time.FixedZone("UTC-5", -5*60*60))},
		},
		{
			[]byte("2018-12-31T09:05:07.949+00:00"),
			Timestamp{Time: time.Date(2018, 12, 31, 9, 5, 7, 949000000, time.UTC)},
		},
	}

	for _, test := range tests {
		t.Run(string(test.in), func(t *testing.T) {
			timestamp := Timestamp{}
			err := timestamp.UnmarshalJSON(test.in)
			if err != nil {
				t.Fatalf("Unexpected error while unmarshaling timestamp: %v", err)
			}

			if !timestamp.Time.Equal(test.out.Time) {
				t.Errorf("got: %s want: %s", timestamp, test.out)
			}
		})
	}
}

func TestErrorError(t *testing.T) {
	want := "not_found: The requested object or REST method was not found."
	err := Error{
		Status:   404,
		Code:     "not_found",
		Details:  "The requested object or REST method was not found.",
		Type:     nil,
		Warnings: []string{},
	}
	if err.Error() != want {
		t.Errorf("got: %s want: %s", err.Error(), want)
	}
}

func TestError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"object": "error", "code": "not_found", "status": 404, "details": "The requested object or REST method was not found."}`)
	})
	client, ts, err := setupTestServer("/cards/nope", handler)
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

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

func TestNewClientWithClientSecret(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "Bearer cs-12345" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, `{"object": "error", "code": "forbidden", "status": 403, "details": ""}`)
			return
		}

		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": []}`)
	})
	client, ts, err := setupTestServer("/symbology", handler, WithClientSecret("cs-12345"))
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	_, err = client.ListCardSymbols(ctx)
	if err != nil {
		t.Fatalf("Error listing card symbols using client with client secret set: %v", err)
	}
}

func TestNewClientWithGrantSecret(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "Bearer 12345" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, `{"object": "error", "code": "forbidden", "status": 403, "details": ""}`)
			return
		}

		fmt.Fprintln(w, `{"object": "list", "has_more": false, "data": []}`)
	})
	client, ts, err := setupTestServer("/symbology", handler, WithGrantSecret("12345"))
	if err != nil {
		t.Fatalf("Error setting up test server: %v", err)
	}
	defer ts.Close()

	ctx := context.Background()
	_, err = client.ListCardSymbols(ctx)
	if err != nil {
		t.Fatalf("Error listing card symbols using client with grant secret set: %v", err)
	}
}

func TestNewClientMultipleSecrets(t *testing.T) {
	_, err := NewClient(WithClientSecret("cs-12345"), WithGrantSecret("12345"))
	if err != ErrMultipleSecrets {
		t.Fatalf("Unexpected error %v received from NewClient when configured with multiple secrets", err)
	}
}
