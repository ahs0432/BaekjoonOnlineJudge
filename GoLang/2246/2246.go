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

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, 2)
		fmt.Fscanln(reader, &table[i][0], &table[i][1])
	}

	count := 0
	for i := 0; i < n; i++ {
		count++
		for j := 0; j < n; j++ {
			if i != j {
				if table[i][0] >= table[j][0] && table[i][1] >= table[j][1] {
					count--
					break
				}
			}
		}
	}
	fmt.Println(count)
}
