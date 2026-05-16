package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var x, y float64
		fmt.Fscanln(reader, &x, &y)

		if x == 0 && y == 0 {
			fmt.Fprintln(writer, "AXIS")
			break
		} else if x == 0 || y == 0 {
			fmt.Fprintln(writer, "AXIS")
		} else if x > 0 && y > 0 {
			fmt.Fprintln(writer, "Q1")
		} else if x < 0 && y < 0 {
			fmt.Fprintln(writer, "Q3")
		} else if x > 0 && y < 0 {
			fmt.Fprintln(writer, "Q4")
		} else {
			fmt.Fprintln(writer, "Q2")
		}
	}

	writer.Flush()
}
