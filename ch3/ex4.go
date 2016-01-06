package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

func xyscale(width int) float64 {
	return float64(width / 2 / xyrange)
}

func zscale(height int) float64 {
	return float64(height) * 0.4
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	color := r.URL.Query().Get("color")
	if len(color) == 0 {
		color = "white"
	}
	height, err := toInt(r.URL.Query().Get("height"), 320)
	if err != nil {
		log.Println(err)
		http.Error(w, "height should be integer", http.StatusBadRequest)
		return
	}
	width, err := toInt(r.URL.Query().Get("width"), 600)
	if err != nil {
		log.Println(err)
		http.Error(w, "width should be integer", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	svg(w, height, width, color)
}

func toInt(param string, d int) (int, error) {
	if len(param) == 0 {
		return d, nil
	}
	h, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return h, nil
}

func svg(w io.Writer, height, width int, color string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(width, height, i+1, j)
			bx, by := corner(width, height, i, j)
			cx, cy := corner(width, height, i, j+1)
			dx, dy := corner(width, height, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(width, height, i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale(width)
	sy := float64(height)/2 + (x+y)*sin30*xyscale(width) - z*zscale(height)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
