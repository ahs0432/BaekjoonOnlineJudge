package main

import (
	"bufio"
	"fmt"
	"os"
)

func divisor(n int) int {
	cnt := 1
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, tmp, divisor(tmp))
	}
	writer.Flush()
}
