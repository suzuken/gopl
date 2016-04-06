package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
)

type Pixel struct {
	x, y int
	z    color.Color
}

// bound is helper func for caliculate adaptive range.
func bound(i, given, all int) (from, to int) {
	if i < 0 {
		return 0, 0
	}
	if i >= all {
		panic("i should be lower than all")
	}
	return i * given / all, (i+1)*given/all - 1
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	var (
		concurrency = flag.Int("concurrency", 2, "go go go (number of goroutine)")
		width       = flag.Int("width", 1024, "width")
		height      = flag.Int("height", 1024, "height")
	)
	flag.Parse()
	var wg sync.WaitGroup
	ch := make(chan Pixel, (*height)*(*width))

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			from, to := bound(i, *height, *concurrency)
			for py := from; py <= to; py++ {
				y := float64(py)/float64(*height)*(ymax-ymin) + ymin
				for px := 0; px < *width; px++ {
					// Image point (px, py) represents complex value z.
					x := float64(px)/float64(*width)*(xmax-xmin) + xmin
					z := complex(x, y)
					ch <- Pixel{px, py, mandelbrot(z)}
				}
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	for p := range ch {
		img.Set(p.x, p.y, p.z)
	}
	if err := png.Encode(os.Stdout, img); err != nil {
		log.Fatalf("encode failed %s", err)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
