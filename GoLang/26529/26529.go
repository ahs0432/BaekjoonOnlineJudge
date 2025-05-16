package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fib := make([]int, 46)

	fib[0], fib[1] = 1, 1

	for i := 2; i < 46; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	var n, tmp int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, fib[tmp])
	}
	writer.Flush()
}
