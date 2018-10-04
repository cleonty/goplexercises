package sexpr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. 3ack D. Ripper":  "Sterling Hayden",
			`Maj. T.3. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	data, err := Marshal(strangelove)
	if err != nil {
		t.Errorf("%v\v", err)
	}
	fmt.Println(string(data))
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Errorf("%v\v", err)
	}
	if !reflect.DeepEqual(strangelove, movie) {
		t.Errorf("after unmarshaling got %v, want %v\v", movie, strangelove)
	}
}
