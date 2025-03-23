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

	sum, up := 0, 0
	t := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t[i])
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})

	for i := 0; i < n && i < 42; i++ {
		sum += t[i]

		if t[i] >= 60 {
			up++
		}

		if t[i] >= 100 {
			up++
		}

		if t[i] >= 140 {
			up++
		}

		if t[i] >= 200 {
			up++
		}

		if t[i] >= 250 {
			up++
		}
	}
	fmt.Println(sum, up)
}
