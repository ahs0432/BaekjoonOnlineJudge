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

	var m, a, b float64
	var t, h, mm, s int
	for {
		fmt.Fscanln(reader, &m, &a, &b)

		if m == 0 && a == 0 && b == 0 {
			break
		}

		t = int(math.Round((m/a - m/b) * 3600.0))
		h = t / 3600
		mm = (t % 3600) / 60
		s = t % 60
		fmt.Fprintf(writer, "%d:%02d:%02d\n", h, mm, s)
	}
	writer.Flush()
}
