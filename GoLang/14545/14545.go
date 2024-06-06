package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp int
	var sum int
	for i := 0; i < n; i++ {
		sum = 0
		fmt.Fscanln(reader, &tmp)

		for j := 1; j <= tmp; j++ {
			sum += j * j
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
