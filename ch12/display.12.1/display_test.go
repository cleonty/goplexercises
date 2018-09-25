package display

import (
	"os"
	"reflect"
	"testing"

	"github.com/cleonty/gopl/ch7/eval"
)

func TestDisplay(t *testing.T) {
	е, _ := eval.Parse("sqrt(A / pi)")
	Display("e", е)
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestDisplayWithMovie(t *testing.T) {
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
	Display("strangelove", strangelove)
}

func TestDisplayWithStderr(t *testing.T) {
	//Display("os.Stderr", os.Stderr)
}

func TestDisplayWithReflectValue(t *testing.T) {
	Display("rV", reflect.ValueOf(os.Stderr))
}

func TestWithArrayAsKey(t *testing.T) {
	m := make(map[[2]int]string)
	m[[...]int{1, 2}] = "leonty"
	m[[...]int{2, 3}] = "opera"
	Display("m", m)
}
