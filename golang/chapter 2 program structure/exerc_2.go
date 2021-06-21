// the main file of the exercises in chapter 2
// pachage name will be chap2Excer
package main

import (
	"fmt"
	"flag"
	"chap2Excer"
)

var prob = flag.Int("p", 1, "Select the number of the exercises.")

func main() {
	flag.Parse()
	swicth *p {
	case 1:
		chap2Excer.solve_2_1()
	default:
		fmt.Println("Invalid exercise number.")
	}
}