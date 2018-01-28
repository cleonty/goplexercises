// main.go
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fractal(w)
}

func fractal(w io.Writer) {
	const (
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height          = 1024, 1024
	)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img)
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128 = z
	for n := uint8(0); n < iterations; n++ {
		v = 1.0 / (v * v * v)
		if math.Abs(cmplx.Abs(v)-1.0) < 0.01 {
			//return color.Gray{255 - contrast*n}
			return color.RGBA{255, 255 - contrast*n, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
