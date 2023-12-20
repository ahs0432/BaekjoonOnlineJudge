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

	var g float64
	var d string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &g, &d)

		if d == "kg" {
			fmt.Fprintf(writer, "%.4f lb\n", g*2.2046)
		} else if d == "lb" {
			fmt.Fprintf(writer, "%.4f kg\n", g*0.4536)
		} else if d == "l" {
			fmt.Fprintf(writer, "%.4f g\n", g*0.2642)
		} else if d == "g" {
			fmt.Fprintf(writer, "%.4f l\n", g*3.7854)
		}
	}

	writer.Flush()
}
