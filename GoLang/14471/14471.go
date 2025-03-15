package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	ans := 0
	now := 0
	t := make([]int, 0)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a >= n {
			now++
		} else {
			t = append(t, n-a)
		}
	}

	if now < m-1 {
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})

		for i := 0; i < (m-1)-now; i++ {
			ans += t[i]
		}
	}
	fmt.Println(ans)
}
