package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func checkCount(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var g string
	fmt.Fscanln(reader, &g)

	var n int
	fmt.Fscanln(reader, &n)
	table := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	dl, do, dv, de := strings.Count(g, "L"), strings.Count(g, "O"), strings.Count(g, "V"), strings.Count(g, "E")
	var l, o, v, e, p int
	max, now := -1, 0

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	for i := 0; i < n; i++ {
		l = dl + strings.Count(table[i], "L")
		o = do + strings.Count(table[i], "O")
		v = dv + strings.Count(table[i], "V")
		e = de + strings.Count(table[i], "E")

		p = ((l + o) * (l + v) * (l + e) * (o + v) * (o + e) * (v + e)) % 100
		if p > max {
			max = p
			now = i
		}
	}
	fmt.Println(table[now])
}
