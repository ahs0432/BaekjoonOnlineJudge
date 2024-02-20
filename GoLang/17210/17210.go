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

	var n int
	var m float64
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	if n <= 5 {
		for i := 1; i < n; i++ {
			m = math.Abs(m - 1)
			fmt.Fprintf(writer, "%0.f\n", m)
		}
	} else {
		fmt.Fprintln(writer, "Love is open door")
	}
	writer.Flush()
}
