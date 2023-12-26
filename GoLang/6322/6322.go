package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	count := 1
	var a, b, c float64
	for {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		fmt.Fprintf(writer, "Triangle #%d\n", count)
		if a == -1 {
			ans := math.Sqrt(c*c - b*b)
			if ans > 0 {
				fmt.Fprintf(writer, "a = %.3f\n", ans)
			} else {
				fmt.Fprintln(writer, "Impossible.")
			}
		} else if b == -1 {
			ans := math.Sqrt(c*c - a*a)
			if ans > 0 {
				fmt.Fprintf(writer, "b = %.3f\n", ans)
			} else {
				fmt.Fprintln(writer, "Impossible.")
			}
		} else {
			ans := math.Sqrt(a*a + b*b)
			if ans > 0 {
				fmt.Fprintf(writer, "c = %.3f\n", ans)
			} else {
				fmt.Fprintln(writer, "Impossible.")
			}
		}
		fmt.Fprintln(writer)
		count++
	}
	writer.Flush()
}
