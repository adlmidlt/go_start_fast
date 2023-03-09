package tutorial_1

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()

		fmt.Printf("Эхо: %s\n", txt)
	}
}
