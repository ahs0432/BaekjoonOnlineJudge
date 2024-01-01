package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var m int
	t := make([]int, 3)
	sum := 0
	for i := 0; i < n; i++ {
		sum = 0
		fmt.Fscanln(reader, &m)
		for j := 0; j < m; j++ {
			fmt.Fscanln(reader, &t[0], &t[1], &t[2])
			sort.Slice(t, func(k, l int) bool {
				return t[k] > t[l]
			})

			if t[0] > 0 {
				sum += t[0]
			}
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
