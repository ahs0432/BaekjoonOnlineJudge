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

	var count, loc, now int
	for i := 1; i < len(table); i++ {
		now = table[i]

		for loc = i - 1; loc >= 0 && now < table[loc]; loc-- {
			table[loc+1] = table[loc]
			count++

			if count == k {
				for j := 0; j < len(table); j++ {
					fmt.Print(table[j], " ")
				}
				fmt.Println()
				return
			}
		}

		if loc+1 != i {
			table[loc+1] = now
			count++
			if count == k {
				for j := 0; j < len(table); j++ {
					fmt.Print(table[j], " ")
				}
				fmt.Println()
				return
			}
		}
	}

	fmt.Println(-1)
}
