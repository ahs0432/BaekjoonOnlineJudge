package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s1, s2, s3 := []int{0, 1, 0, 0}, []int{0, 0, 1, 0}, []int{0, 0, 0, 1}
	a1, a2, a3 := 0, 0, 0

	var n int
	fmt.Fscanln(reader, &n)

	var a, b, g int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &g)

		s1[a], s1[b] = s1[b], s1[a]
		s2[a], s2[b] = s2[b], s2[a]
		s3[a], s3[b] = s3[b], s3[a]

		if s1[g] == 1 {
			a1++
		}
		if s2[g] == 1 {
			a2++
		}
		if s3[g] == 1 {
			a3++
		}
	}

	ans := []int{a1, a2, a3}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	fmt.Println(ans[0])
}
