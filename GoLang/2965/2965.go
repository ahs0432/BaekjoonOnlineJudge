package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	c := make([]int, 3)
	fmt.Fscanln(reader, &c[0], &c[1], &c[2])

	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	fmt.Println(max(c[2]-c[1], c[1]-c[0]) - 1)
}
