// exercise 2.1 page 42
package chap2Excer

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const {
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	ZeroK Kelvin = -273.15
}

func (c Celsius) String() string {return fmt.Sprintf("%v°C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%v°F", f)}
func (k Kelvin) String() string {return fmt.Sprintf("%v°K", k)}

func cToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5+32)}
func fToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}

func kToC(k Kelvin) Celsius {return Celsius(k-ZeroK)}
func cToK(c Celsius) Kelvin {return Kelvin(c+ZeroK)}

func kToF(k Kelvin) Fahrenheit {return cToF(kToC)}
func fToK(f Fahrenheit) Kelvin {return cToK(fToC)}

func solve_2_1() {
	fmt.Println("Convert K to C: ", kToC(Kelvin(1)))
	fmt.Println("Convert C to K: ", cToK(Celsius(1)))
	fmt.Println("Convert K to F: ", kToF(Kelvin(1)))
	fmt.Println("Convert F to K: ", kToC(Fahrenheit(1)))
}