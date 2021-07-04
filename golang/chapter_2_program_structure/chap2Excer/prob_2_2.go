// exercise 2.2
package chap2Excer

func Feet2Meter(f Feet) Meter {return Meter(f*Feet(FtoM))}
func Meter2Feet(m Meter) Feet {return Feet(m/Meter(FtoM))}

func Pounds2Kilo(p Pounds) Kilograms {return Kilograms(p*Pounds(PtoKg))}
func Kilo2Pounds(k Kilograms) Pounds {return Pounds(k/Kilograms(PtoKg))}