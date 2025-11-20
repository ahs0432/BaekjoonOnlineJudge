package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	count := 0
	var key int
	for i := 1; i < n; i++ {
		key = table[i]

		j := i - 1
		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			count++

			if count == k {
				fmt.Println(table[j])
				return
			}
			j--
		}

		if table[j+1] != key {
			count++
			table[j+1] = key
		}

		if count == k {
			fmt.Println(key)
			return
		}
	}

	fmt.Println(-1)
}
