package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := make([]int, 3)
	fmt.Fscanln(reader, &a[0], &a[1], &a[2])

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	var s string
	fmt.Fscanln(reader, &s)

	for _, c := range s {
		fmt.Print(a[c-65], " ")
	}
	fmt.Println()
}
