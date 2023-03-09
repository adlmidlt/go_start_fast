package tutorial_2

import (
	"fmt"
	"io"
	"os"
)

func Stdouterr() {
	io.WriteString(os.Stdout, "Это строка для стандартного вывода.\n")
	io.WriteString(os.Stderr, "Это строка для стандартного вывода ошибки.\n")

	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, err := os.Stdout.Write(buf); err != nil {
			panic(err)
		}
	}

	fmt.Fprintln(os.Stdout, "\n")
}
