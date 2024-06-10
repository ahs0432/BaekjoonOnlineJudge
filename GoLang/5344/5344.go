package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	tmp := 0
	if a > b {
		a, b = b, a
	}

	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		fmt.Fprintln(writer, gcd(a, b))
	}
	writer.Flush()
}
