package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	whiteIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100
		nframes = 64
		delay   = 8
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
	)
	green := color.RGBA{0x00, 0x80, 0x00, 0xff}
	red := color.RGBA{0xff, 0x00, 0x00, 0xff}
	blue := color.RGBA{0xff, 0x00, 0xff, 0xff}
	palette := []color.Color{color.White, green, red, blue}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cycles int
	c := r.URL.Query().Get("cycles")
	if len(c) == 0 {
		cycles = 5
	} else {
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Println(err)
		}
		cycles = i
	}
	lissajous(w, cycles)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
