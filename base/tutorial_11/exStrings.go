package tutorial_11

import "fmt"

func MainString() {
	//printfString()
	printByByteShalom()
}

func printfString() {
	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}

func printByByteShalom() {
	str := "shalom"
	s := str[0]
	h := str[1]
	a := str[2]
	l := str[3]
	o := str[4]
	m := str[5]
	fmt.Printf("%c\n%c\n%c\n%c\n%c\n%c\n", s, h, a, l, o, m)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}
