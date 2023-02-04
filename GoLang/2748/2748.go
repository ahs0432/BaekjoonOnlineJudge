package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	if n == 0 {
		fmt.Println(0)
		return
	} else if n <= 2 {
		fmt.Println(1)
		return
	}

	table := make([]int64, n+1)
	table[1] = 1
	table[2] = 1

	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	fmt.Println(table[n])
}
