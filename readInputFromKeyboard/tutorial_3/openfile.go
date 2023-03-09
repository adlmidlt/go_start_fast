package tutorial_3

import (
	"fmt"
	"io"
	"os"
)

func Openfile() {
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("temp/file.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	io.WriteString(f, "Test String")
	f.Close()
}
