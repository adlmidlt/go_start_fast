package tutorial_7

import (
	"fmt"
	"math/rand"
	"time"
)

func BuyTicket() {
	flyingToMars()
}

const (
	hoursPerDay             = 24
	distanceFromEarthToMars = 62100000
)

func flyingToMars() {
	companySpace := ""
	trip := ""

	rand.Seed(time.Now().UnixNano())
	fmt.Println("SpaceLine         Days   Trip type   Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			companySpace = "Virgin Galactic"
		case 1:
			companySpace = "SpaceX"
		case 2:
			companySpace = "Space Adventure"
		}

		speedShip := rand.Intn(15) + 16
		durationOfTrip := distanceFromEarthToMars / speedShip / hoursPerDay
		priceTicket := rand.Intn(35) + 16

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			priceTicket *= 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %7v %-11v $%4v\n", companySpace, durationOfTrip, trip, priceTicket)
	}
}
