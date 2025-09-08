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

	t := make([]int, 3)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &t[0], &t[1], &t[2])
		sort.Slice(t, func(j, k int) bool {
			return t[j] < t[k]
		})

		fmt.Fprintf(writer, "Case #%d: ", i)
		if t[0]+t[1] <= t[2] {
			fmt.Fprintln(writer, "invalid!")
		} else if t[0] == t[1] && t[1] == t[2] {
			fmt.Fprintln(writer, "equilateral")
		} else if t[0] == t[1] || t[1] == t[2] || t[0] == t[2] {
			fmt.Fprintln(writer, "isosceles")
		} else {
			fmt.Fprintln(writer, "scalene")
		}
	}

	writer.Flush()
}
