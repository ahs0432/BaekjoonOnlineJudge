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

	for i := 1; i <= n; i++ {
		fmt.Fprint(writer, i, " ")

		if n == i || i%6 == 0 {
			fmt.Fprint(writer, "Go! ")
		}
	}
	writer.Flush()
}
