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
	for i := 0; i < n-1; i++ {
		if table[i] < table[i+1] {
			continue
		} else {
			table[i+1] += k
			count++
		}
	}

	check := false
	for i := 0; i < n-1; i++ {
		if table[i] >= table[i+1] {
			check = true
			break
		}
	}

	if check {
		fmt.Println(-1)
	} else {
		fmt.Println(count)
	}
}
