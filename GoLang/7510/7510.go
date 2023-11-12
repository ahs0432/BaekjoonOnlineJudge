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

		sort.Slice(t, func(x, y int) bool {
			return t[x] < t[y]
		})

		if t[0]*t[0]+t[1]*t[1] == t[2]*t[2] {
			fmt.Fprintf(writer, "Scenario #%d:\nyes\n\n", i)
		} else {
			fmt.Fprintf(writer, "Scenario #%d:\nno\n\n", i)
		}
	}
	writer.Flush()
}
