package scryfall

import (
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
