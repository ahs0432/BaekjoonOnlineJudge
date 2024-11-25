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

	var r, w, l int
	var tmp float64
	for i := 1; ; i++ {
		fmt.Fscanln(reader, &r, &w, &l)
		if r == 0 {
			break
		}
		tmp = math.Sqrt(float64(w*w + l*l))
		if float64(r*2) >= tmp {
			fmt.Fprintln(writer, "Pizza", i, "fits on the table.")
		} else {
			fmt.Fprintln(writer, "Pizza", i, "does not fit on the table.")
		}
	}
	writer.Flush()
}
