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

	a := make([]int, 100)
	a[1] = 1
	for i := 3; i < 100; i += 2 {
		a[i] += a[i-2] + i
	}

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, a[tmp])
	}
	writer.Flush()
}
