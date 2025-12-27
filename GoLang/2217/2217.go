package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	rope := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &rope[i])
	}
	sort.Slice(rope, func(i, j int) bool {
		return rope[i] < rope[j]
	})

	max := rope[n-1]
	for i := n - 1; i >= 0; i-- {
		current := rope[i] * (n - i)
		if max < current {
			max = current
		}
	}
	fmt.Println(max)
}
