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

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	count := 1
	for i := 1; i < n; i++ {
		if (table[i-1]+table[i])%2 == 1 {
			count++
		}
	}
	fmt.Println(count)
}
