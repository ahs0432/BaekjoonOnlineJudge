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

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp&(-tmp) == tmp {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
	writer.Flush()
}
