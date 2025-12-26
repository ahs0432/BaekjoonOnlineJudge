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

	a := make([]int, n)
	b := make([]int, n)

	sum := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &b[i])
		} else {
			fmt.Fscan(reader, &b[i])
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	for i := 0; i < n; i++ {
		sum += (a[i] * b[i])
	}
	fmt.Println(sum)
}
