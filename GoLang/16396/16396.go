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

	table := make([]bool, 10001)
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)
		for j := x; j < y; j++ {
			table[j] = true
		}
	}

	count := 0
	for i := 1; i <= 10000; i++ {
		if table[i] {
			count++
		}
	}
	fmt.Println(count)
}
