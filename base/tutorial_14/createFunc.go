package tutorial_14

import "fmt"

func CreateFunc() {
	result()
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15

	return k
}

func result() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)

	fmt.Printf("%.2f K is %.2f C", kelvin, celsius)
}
