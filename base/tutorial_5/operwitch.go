package tutorial_5

import (
	"fmt"
	"time"
)

func OperatorSwitch() {
	//weekday()
	//exampleSwitch()
	//breakSwitch()
	switchType()
}

func weekday() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Сегодня понедельник.")
	case time.Tuesday:
		fmt.Println("Сегодня вторник.")
	case time.Wednesday:
		fmt.Println("Сегодня среда.")
	case time.Thursday:
		fmt.Println("Сегодня четверг.")
	case time.Friday:
		fmt.Println("Сегодня пятница.")
	case time.Sunday:
		fmt.Println("Сегодня суббота.")
	case time.Saturday:
		fmt.Println("Сегодня воскресенье.")
	default:
		fmt.Println("Неизвестный день недели.")
	}
}

func exampleSwitch() {
	switch num := 6; num%2 == 0 {
	case true:
		fmt.Println("even day")

	case false:
		fmt.Println("odd value")
	}
}

func breakSwitch() {
	w := "a b c\td\nefg hi"

	for _, e := range w {
		switch e {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c\n", e)
		}
	}
}

func switchType() {
	var data interface{}
	data = 112523652346.23463246345

	switch myType := data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("%T", myType)
	}
}
