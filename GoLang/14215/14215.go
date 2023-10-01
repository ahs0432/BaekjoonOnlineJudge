package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := make([]int, 3)
	fmt.Fscanln(reader, &a[0], &a[1], &a[2])
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	fmt.Println(a[0] + a[1] + min(a[2], a[0]+a[1]-1))
}
