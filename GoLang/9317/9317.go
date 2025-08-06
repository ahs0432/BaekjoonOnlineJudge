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

	var d, h, v float64
	var weight, height float64
	for {
		fmt.Fscanln(reader, &d, &h, &v)

		if d == 0 && h == 0 && v == 0 {
			break
		}

		weight = 16 * d / math.Sqrt(337)
		height = 9 * weight / 16

		fmt.Fprintf(writer, "Horizontal DPI: %.2f\n", h/weight)
		fmt.Fprintf(writer, "Vertical DPI: %.2f\n", v/height)
	}
	writer.Flush()
}
