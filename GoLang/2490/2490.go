package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < 3; i++ {
		a := 0
		var tmp int
		for j := 0; j < 4; j++ {
			fmt.Fscan(reader, &tmp)
			a += tmp
		}

		if a == 0 {
			fmt.Fprintln(writer, "D")
		} else if a == 1 {
			fmt.Fprintln(writer, "C")
		} else if a == 2 {
			fmt.Fprintln(writer, "B")
		} else if a == 3 {
			fmt.Fprintln(writer, "A")
		} else if a == 4 {
			fmt.Fprintln(writer, "E")
		}
	}
	writer.Flush()
}
