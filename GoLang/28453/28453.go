package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)

		if tmp == 300 {
			fmt.Fprint(writer, "1 ")
		} else if tmp >= 275 {
			fmt.Fprint(writer, "2 ")
		} else if tmp >= 250 {
			fmt.Fprint(writer, "3 ")
		} else {
			fmt.Fprint(writer, "4 ")
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
