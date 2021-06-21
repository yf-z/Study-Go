// exercise 2.1 page 42
package chap2Excer

import "fmt"

func cToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5+32)}
func fToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}

func kToC(k Kelvin) Celsius {return Celsius(k-ZeroK)}
func cToK(c Celsius) Kelvin {return Kelvin(c+Celsius(ZeroK))}

func kToF(k Kelvin) Fahrenheit {return cToF(kToC(k))}
func fToK(f Fahrenheit) Kelvin {return cToK(fToC(f))}

func Solve_2_1() {
	fmt.Println("Convert K to C: ", kToC(Kelvin(1)))
	fmt.Println("Convert C to K: ", cToK(Celsius(1)))
	fmt.Println("Convert K to F: ", kToF(Kelvin(1)))
	fmt.Println("Convert F to K: ", fToK(Fahrenheit(1)))
}