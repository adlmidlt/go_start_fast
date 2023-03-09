package tutorial_19

import "fmt"

func SliceArr() {

	planets := []string{"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун"}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()

	fmt.Println(planets)
}

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "Новый " + p[i]
	}
}
