package tutorial_8

import (
	"fmt"
	"math/rand"
	"time"
)

func Float3264() {
	countMonet := 0.0
	maxMonet := 20.0

	for {
		rand.Seed(time.Now().UnixNano())
		if countMonet <= maxMonet {
			switch rand.Intn(2) + 1 {
			case 1:
				countMonet += 0.05
			case 2:
				countMonet += 0.10
			case 3:
				countMonet += 0.25
			}
			fmt.Printf("Кол-во монет: %.2f\n", countMonet)
		} else {
			break
		}
	}
}
