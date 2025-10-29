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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	check := true
	for i := 1; i <= n; i++ {
		if i != a[i-1] {
			fmt.Println(i)
			check = false
			break
		}
	}

	if check {
		fmt.Println(n + 1)
	}
}
