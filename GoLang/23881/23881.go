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

	num := 0
	var tmp, max, a, b, count int
	for i := n - 1; i >= 0; i-- {
		tmp = table[0]
		max = 0

		for j := 0; j < i; j++ {
			if tmp < table[j] {
				tmp = table[j]
				max = j
			}
		}

		if table[max] > table[i] {
			count = table[i]
			table[i] = table[max]
			table[max] = count
			a = table[max]
			b = table[i]
			num++

			if num == k {
				fmt.Println(a, b)
				return
			}
		}
	}

	fmt.Println(-1)
}
