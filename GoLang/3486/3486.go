package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(a int) int {
	new := 0
	for a != 0 {
		new *= 10
		new += (a % 10)
		a /= 10
	}
	return new
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, reverse(reverse(a)+reverse(b)))
	}
	writer.Flush()
}
