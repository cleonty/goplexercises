package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/cleonty/gopl/ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("пустое выражение")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("Неизвестная переменная %q", v)
		}
	}
	return expr, nil
}

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(w io.Writer, f func(x, у float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if !math.IsNaN(ax) && !math.IsNaN(ay) &&
				!math.IsNaN(bx) && !math.IsNaN(by) &&
				!math.IsNaN(cx) && !math.IsNaN(cy) &&
				!math.IsNaN(dx) && !math.IsNaN(dy) {
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(os.Stdout, "skip <polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, у float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Sin(y))
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "некорректное выражение: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // Расстояние от (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	http.HandleFunc("/", plot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
