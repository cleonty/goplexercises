package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/gif")
	fractal(w)
}

func fractal(out io.Writer) {
	const (
		size    = 400
		nframes = 5
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	anim := gif.GIF{LoopCount: nframes}
	for i := 0; i < nframes; i++ {
		var a, b, c, d, e, f float64
		var x, y float64
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		npoints := (nframes + 1) * 5 * 70000
		for i := 0; i < npoints; i++ {
			r := rand.Float32()
			if r > 0 && r < 0.01 {
				a = 0
				b = 0
				c = 0
				d = 0.16
				e = 0
				f = 0
			} else if r > 0.01 && r < 0.8 {
				a = 0.85
				b = 0.04
				c = -0.04
				d = 0.85
				e = 0
				f = 1.6
			} else if r > 0.8 && r < 0.9 {
				a = 0.2
				b = -0.26
				c = 0.23
				d = 0.22
				e = 0
				f = 1.6
			} else if r > 0.8 && r < 1.0 {
				a = -0.15
				b = 0.28
				c = 0.26
				d = 0.24
				e = 0
				f = 0.44
			}
			x1 := (a * x) + (b * y) + e
			y1 := (c * x) + (d * y) + f
			x = x1
			y = y1
			xc1 := size + x*size/3.0
			yc1 := 2*size - y*size*2.0/12.0
			xc := xc1
			yc := yc1
			img.SetColorIndex(int(xc), int(yc), greenIndex)
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
