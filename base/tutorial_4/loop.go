package tutorial_4

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Loop() {
	//trueFalse()
	//equality()
	//ifElse()
	//isLeap()
	//loopFor()
	chance()
}

func trueFalse() {
	fmt.Println("Вы находитесь в тёмной пищере.")

	var command = "выйду наружу."
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пищеру:", exit)
}

func equality() {
	fmt.Println("На знаке снаружи написано 'Несовершеннолетним вход запрещен'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("В возрасте %v, я совершеннолетний? %v\n", age, minor)
}

func ifElse() {
	var command = "идти на восток"

	if command == "идти на восток" {
		fmt.Println("Вы направляетесь к горе.")
	} else if command == "зайти внутрь" {
		fmt.Println("Вы заходите в пищеру, где будете жить до конца своих дней.")
	} else {
		fmt.Println("Пока не совсем понятно.")
	}
}

func isLeap() {
	fmt.Println("На дворе 2100 год. Он високосный?")

	var year = 2000
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Этот год високосный!")
	} else {
		fmt.Println("К сожалению, нет. Этот год не високосный.")
	}
}

func loopSwitch() {
	var command = "идти на восток"

	switch command {
	case "идти на восток":
		fmt.Println("Вы направляетесь к горе.")
	case "зайти внутрь":
		fmt.Println("Вы заходите в пищеру, где будете жить до конца своих дней.")
	default:
		fmt.Println("Пока не совсем понятно.")
	}
}

func loopFor() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Запуск!")
}

func chance() {
	fmt.Println("Загайдай число")
	var hiddenNum = 39

	for {
		var randNum = rand.Intn(50) + 1
		if randNum > hiddenNum {
			fmt.Println("Загаданное число меньше ===>", randNum)
		} else if randNum < hiddenNum {
			fmt.Println("Загадочное число больше ===>", randNum)
		} else {
			fmt.Println("Вы угадали число!!! ===>", randNum)
			break
		}
	}
}
