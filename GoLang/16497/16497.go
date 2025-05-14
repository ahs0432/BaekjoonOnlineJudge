package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a, b, k int
	fmt.Fscanln(reader, &n)

	table := make([]int, 31)

	max := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		for ; a < b; a++ {
			table[a] += 1
			if table[a] > max {
				max = table[a]
			}
		}
	}

	fmt.Fscanln(reader, &k)

	if max <= k {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
