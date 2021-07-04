// exercise 2.1 page 42
package chap2Excer

import "fmt"

func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5+32)}
func FToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}

func KToC(k Kelvin) Celsius {return Celsius(k-ZeroK)}
func CToK(c Celsius) Kelvin {return Kelvin(c+Celsius(ZeroK))}

func KToF(k Kelvin) Fahrenheit {return CToF(KToC(k))}
func FToK(f Fahrenheit) Kelvin {return CToK(FToC(f))}

func Solve_2_1() {
	fmt.Println("Convert K to C: ", KToC(Kelvin(1)))
	fmt.Println("Convert C to K: ", CToK(Celsius(1)))
	fmt.Println("Convert K to F: ", KToF(Kelvin(1)))
	fmt.Println("Convert F to K: ", FToK(Fahrenheit(1)))
}