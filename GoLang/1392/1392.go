package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	nl := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &nl[i])
	}

	var tmp int
	var sum int
	for i := 0; i < q; i++ {
		fmt.Fscanln(reader, &tmp)

		sum = 0
		for j := 0; j < n; j++ {
			sum += nl[j]
			if tmp < sum {
				fmt.Fprintln(writer, j+1)
				break
			}
		}
	}
	writer.Flush()
}
