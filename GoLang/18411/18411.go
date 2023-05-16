package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := make([]int, 3)
	fmt.Fscanln(reader, &n[0], &n[1], &n[2])

	sort.Slice(n, func(i, j int) bool {
		return n[i] > n[j]
	})

	fmt.Println(n[0] + n[1])
}
