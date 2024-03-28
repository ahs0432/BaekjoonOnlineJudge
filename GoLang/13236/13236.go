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

	fmt.Fprint(writer, n, " ")
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			fmt.Fprint(writer, n, " ")
		} else {
			n = n*3 + 1
			fmt.Fprint(writer, n, " ")
		}
	}
	writer.Flush()
}
