package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, v int
	fmt.Fscanln(reader, &n, &m, &v)

	table := make([][]int, n)
	for i := 0; i < m; i++ {
		var s, d int
		fmt.Fscanln(reader, &s, &d)
		s--
		d--
		table[s] = append(table[s], d)
		table[d] = append(table[d], s)
	}

	visit := make([]bool, n)

	q := []int{v - 1}
	dfs := []int{}

	for len(q) != 0 {
		now := q[0]
		q = q[1:]

		if visit[now] {
			continue
		}

		visit[now] = true
		dfs = append(dfs, now+1)
		sort.Slice(table[now], func(i, j int) bool {
			return table[now][i] < table[now][j]
		})

		qtmp := []int{}
		for _, t := range table[now] {
			if !visit[t] {
				qtmp = append(qtmp, t)
			}
		}
		q = append(qtmp, q...)
	}

	visit = make([]bool, n)
	visit[v-1] = true

	q = []int{v - 1}
	bfs := []int{v}
	for len(q) != 0 {
		now := q[0]
		q = q[1:]

		for _, t := range table[now] {
			if !visit[t] {
				q = append(q, t)
				visit[t] = true
				bfs = append(bfs, t+1)
			}
		}
	}

	for _, d := range dfs {
		fmt.Print(d, " ")
	}
	fmt.Println()
	for _, b := range bfs {
		fmt.Print(b, " ")
	}
	fmt.Println()
}
