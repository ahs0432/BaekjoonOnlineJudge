package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, k, sum, tmp int
	for i := 0; i < t; i++ {
		sum = 0
		fmt.Fscan(reader, &n, &k)

		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &tmp)
			sum += tmp / k
		}

		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
