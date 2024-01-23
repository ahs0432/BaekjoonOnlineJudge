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

	var b, n, i float64
	var tmp1, tmp2 float64
	for {
		fmt.Fscanln(reader, &b, &n)
		if b == 0 && n == 0 {
			break
		}

		i = 0
		for {
			if math.Pow(i, n) <= b {
				tmp1 = i
			} else if math.Pow(i, n) > b {
				tmp2 = i
				break
			}
			i++
		}

		if math.Abs(math.Pow(tmp1, n)-b) < math.Abs(math.Pow(tmp2, n)-b) {
			fmt.Fprintf(writer, "%0.f\n", tmp1)
		} else {
			fmt.Fprintf(writer, "%0.f\n", tmp2)
		}
	}
	writer.Flush()
}
