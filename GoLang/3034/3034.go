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

	var n, w, h int
	fmt.Fscanln(reader, &n, &w, &h)

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp <= w || tmp <= h || tmp <= int(math.Sqrt(float64(w*w+h*h))) {
			fmt.Fprintln(writer, "DA")
		} else {
			fmt.Fprintln(writer, "NE")
		}
	}
	writer.Flush()
}
