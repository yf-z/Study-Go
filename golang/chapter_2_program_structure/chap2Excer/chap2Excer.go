package chap2Excer

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
// feet meter conversion
type Feet float64
type Meter float64
// pounds kilograms conversion
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	ZeroK Kelvin = -273.15
	// feet meter conversion
	FtoM float64 = 0.3048
	// pounds kilo conversion
	PtoKg float64 = 0.45359237
)

func (c Celsius) String() string {return fmt.Sprintf("%g°C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%g°F", f)}
func (k Kelvin) String() string {return fmt.Sprintf("%g°K", k)}
// feet meter conversion
func (f Feet) String() string {return fmt.Sprintf("%g'", f)}
func (m Meter) String() string {return fmt.Sprintf("%gm", m)}
// pounds kilograms conversion
func (p Pounds) String() string {return fmt.Sprintf("%glb", p)}
func (k Kilograms) String() string {return fmt.Sprintf("%gkg", k)}