package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	_ "os"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}
		cycles, err := strconv.Atoi(r.Form.Get("cycles"))
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		size, err := strconv.Atoi(r.Form.Get("size"))
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		lissajous(w, cycles, size)
	}
	http.HandleFunc("/", handler)
	//lissajous(os.Stdout)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajous(out io.Writer, cycles, size int) {
	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)
	if cycles < 1 {
		cycles = 1
	}
	if size < 10 {
		size = 10
	}
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*math.Pi*float64(cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.01
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
