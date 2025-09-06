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

	h := table[0]
	max := 0
	count := 0

	for i := 1; i < n; i++ {
		if table[i] < h {
			count += 1
		} else {
			h = table[i]
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}
	fmt.Println(max)
}
