package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y, tmp int
	fmt.Fscanln(reader, &x, &y)

	for i := 1; i <= x*y; i++ {
		tmp = 0

		if i%x == 0 {
			tmp += 2
		}

		if i%y == 0 {
			tmp += 1
		}

		if tmp != 0 {
			fmt.Fprint(writer, tmp)
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
