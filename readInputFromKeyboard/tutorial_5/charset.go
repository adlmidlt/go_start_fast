package tutorial_5

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

func Charset() {
	encoder := charmap.Windows1251.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}
	os.WriteFile("example.txt", []byte(s), os.ModePerm)

	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
