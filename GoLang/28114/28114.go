package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	p := make([][]int, 3)
	y := make([]int, 3)
	n := make([]string, 3)

	for i := 0; i < 3; i++ {
		p[i] = make([]int, 2)
		p[i][0] = i
		fmt.Fscanln(reader, &p[i][1], &y[i], &n[i])
	}

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	sort.Slice(p, func(i, j int) bool {
		return p[i][1] > p[j][1]
	})

	for i := 0; i < len(y); i++ {
		fmt.Print(y[i] % 100)
	}
	fmt.Println()

	for i := 0; i < len(p); i++ {
		fmt.Print(string(n[p[i][0]][0]))
	}
	fmt.Println()
}
