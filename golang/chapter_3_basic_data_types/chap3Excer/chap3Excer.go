// all consts and shared values/types
package chap3Excer

import "math"

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width/2/xyrange
	zsclae = height*0.4
	angle = math.Pi/6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)