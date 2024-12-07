package main

import (
	"bufio"
	"fmt"
	"os"
)

func fac(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(writer, "n e")
	fmt.Fprintln(writer, "- -----------")

	var sum float64 = 0.0
	for i := 0; i < 10; i++ {
		sum = 1/fac(float64(i)) + sum

		if i <= 1 {
			fmt.Fprintf(writer, "%d %.0f\n", i, sum)
		} else if i == 2 {
			fmt.Fprintf(writer, "%d %.1f\n", i, sum)
		} else {
			fmt.Fprintf(writer, "%d %.9f\n", i, sum)
		}
	}
	writer.Flush()
}
