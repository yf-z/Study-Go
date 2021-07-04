// the main file of the exercises in chapter 2
// pachage name will be chap2Excer
package main

import (
	"fmt"
	"flag"

	chap2Excer "github.com/yf-z/Study-Go/golang/chapter_2_program_structure/chap2Excer"
)

var prob = flag.Int("p", 1, "Select the number of the exercises.")
var conv = flag.String("c", "", "Select from KC, FC, FK, FM, KP")
var val = flag.Float64("v", 0.0, "Type in the value")

func main() {
	flag.Parse()
	switch *prob {
	case 1:
		chap2Excer.Solve_2_1()
	case 2:
		switch *conv {
		case "KC":
			fmt.Println(chap2Excer.Kelvin(*val), " = ", chap2Excer.KToC(chap2Excer.Kelvin(*val)))
			fmt.Println(chap2Excer.Celsius(*val), " = ", chap2Excer.CToK(chap2Excer.Celsius(*val)))
		case "FC":
			fmt.Println(chap2Excer.Fahrenheit(*val), " = ", chap2Excer.FToC(chap2Excer.Fahrenheit(*val)))
			fmt.Println(chap2Excer.Celsius(*val), " = ", chap2Excer.CToF(chap2Excer.Celsius(*val)))
		case "FK":
			fmt.Println(chap2Excer.Fahrenheit(*val), " = ", chap2Excer.FToK(chap2Excer.Fahrenheit(*val)))
			fmt.Println(chap2Excer.Kelvin(*val), " = ", chap2Excer.KToF(chap2Excer.Kelvin(*val)))
		case "FM":
			fmt.Println(chap2Excer.Feet(*val), " = ", chap2Excer.Feet2Meter(chap2Excer.Feet(*val)))
			fmt.Println(chap2Excer.Meter(*val), " = ", chap2Excer.Meter2Feet(chap2Excer.Meter(*val)))
		case "KP":
			fmt.Println(chap2Excer.Pounds(*val), " = ", chap2Excer.Pounds2Kilo(chap2Excer.Pounds(*val)))
			fmt.Println(chap2Excer.Kilograms(*val), " = ", chap2Excer.Kilo2Pounds(chap2Excer.Kilograms(*val)))
		default:
			fmt.Println("Invalid conversion type.")
		}
	default:
		fmt.Println("Invalid exercise number.")
	}
}