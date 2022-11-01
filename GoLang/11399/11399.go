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

	t := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &t[i])
		} else {
			fmt.Fscan(reader, &t[i])
		}
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	num := 0
	total := 0

	for _, i := range t {
		num += i
		total += (num)
	}

	fmt.Println(total)
}
