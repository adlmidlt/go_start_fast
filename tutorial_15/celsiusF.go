package tutorial_15

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 274.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 0.5 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.0)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func MainStart() {
	var k kelvin = 294.0
	c := k.celsius()
	b := k.fahrenheit()
	fmt.Print(c, "   ", b)
}
