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

	fib := []int64{1, 1, 2, 4}

	for i := 4; i < 69; i++ {
		fib = append(fib, fib[i-1]+fib[i-2]+fib[i-3]+fib[i-4])
	}

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, fib[tmp])
	}
	writer.Flush()
}
