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

	var n float64
	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}

		fmt.Fprintf(writer, "%.2f\n", 1+n+math.Pow(n, 2)+math.Pow(n, 3)+math.Pow(n, 4))
	}
	writer.Flush()
}
