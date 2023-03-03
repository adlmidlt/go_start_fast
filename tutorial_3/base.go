package tutorial_3

import (
	"fmt"
	"math/rand"
)

func Base() {
	//calcWeightLossOnMars()
	//calcOfTimeForTravelToMars()
	//randNumber()
}

// calcWeightLossOnMars - расчёт потери веса на марсе.
func calcWeightLossOnMars() {
	fmt.Println("Мой вес на поверхности Марса равен", 55.0*0.3783, "килограмм, а мой возраст равен", 41*365/687, "годам.")
	fmt.Printf("Мой вес на поверхности Марса равен %v килограмм, а мой возраст равен %v годам.\n", 55.0*0.3783, 41*365/687)
	fmt.Printf("Мой вес на поверхности Марса равен %-10v килограмм, а мой возраст равен %v годам.\n", 55.0*0.3783, 41*365/687)
}

// calcOfTimeForTravelToMars - расчёт времени для путешествия на марс.
func calcOfTimeForTravelToMars() {
	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "секунд.")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "секунд.")

	const hoursPerDay = 24
	var (
		spacexSpeed                 = 100800 // км/ч.
		distanceBetweenMarsAndEarth = 96300000
	)

	fmt.Println(distanceBetweenMarsAndEarth/spacexSpeed/hoursPerDay, "дней.")
}

func randNumber() {
	var randNum = rand.Intn(10) + 1
	fmt.Println(randNum)
}
