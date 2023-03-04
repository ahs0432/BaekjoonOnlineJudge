package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	table := make([]int64, n+1)
	table[0] = 1

	for i := int64(1); i <= n; i++ {
		table[i] = table[i-1] * i
	}

	fmt.Println(table[n])
}
