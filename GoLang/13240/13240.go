package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y int
	fmt.Fscanln(reader, &y, &x)

	line := make([]string, 2)

	for i := 0; i < len(line); i++ {
		for j := 0; j < x; j++ {
			if (j+i)%2 == 0 {
				line[i] += "*"
			} else {
				line[i] += "."
			}
		}
	}

	for i := 0; i < y; i++ {
		fmt.Fprintln(writer, line[i%2])
	}
	writer.Flush()
}
