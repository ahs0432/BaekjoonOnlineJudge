package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(n, j int) int {
	sum := 0
	for n != 0 {
		if n%j == 0 {
			sum++
		} else {
			break
		}
		n /= j
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, sum int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		sum = 0
		for j := 2; j <= n; j++ {
			sum += f(n, j)
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
