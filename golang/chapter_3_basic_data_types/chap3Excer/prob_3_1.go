// excercise 3.1
package chap3Excer

import (
	"fmt"
	"math"
)

func solve_3_1() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
				"style='stroke: grey; fill: white; stroke-width:0.7' "+
				"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<Polygon points = '%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, cy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange*(float64(i)/cells - 0.5)
	y := xyrange*(float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2+(x-y)*cos30*xyscale
	sy := height/2+(x+y)*sin30*xyscale-z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	// check if r is a valid float64
	if math.IsNaN(math.Sin(r)/r) {
		return r
	}
	else {
		return math.Sin(r)/r
	}
}