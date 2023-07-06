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
		fmt.Fscanln(reader, &tmp)

		if tmp%2 == 0 {
			fmt.Fprintln(writer, tmp, "is even")
		} else {
			fmt.Fprintln(writer, tmp, "is odd")
		}
	}
	writer.Flush()
}
