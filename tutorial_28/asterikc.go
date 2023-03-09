package tutorial_28

import "fmt"

func Asterics() {
	asterAper()
}

func asterAper() {
	answer := 23
	fmt.Println(*&answer)

	address := &answer
	fmt.Println(*address)
}
