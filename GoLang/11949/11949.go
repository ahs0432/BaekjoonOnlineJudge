package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < n-1; j++ {
			if table[j]%i > table[j+1]%i {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(table[i])
	}
}
