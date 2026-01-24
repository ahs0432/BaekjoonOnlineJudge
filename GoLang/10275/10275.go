package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(x int64) int64 {
	count := int64(0)
	for x != 0 {
		count += x % 2
		x /= 2
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, a, b int64
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &a, &b)
		fmt.Fprintln(writer, check(int64(math.Pow(2.0, float64(n)))^a^b))
	}
	writer.Flush()
}
