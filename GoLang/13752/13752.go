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

		s := ""
		for j := 0; j < tmp; j++ {
			s += "="
		}

		fmt.Fprintln(writer, s)
	}

	writer.Flush()
}
