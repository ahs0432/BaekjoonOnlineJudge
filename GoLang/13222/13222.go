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
	var w, h float64
	fmt.Fscanln(reader, &n, &w, &h)

	dia := int(math.Sqrt(w*w + h*h))

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		if tmp <= dia {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
