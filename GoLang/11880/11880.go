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

	s := make([]int, 3)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s[0], &s[1], &s[2])

		sort.Slice(s, func(a, b int) bool {
			return s[a] < s[b]
		})

		fmt.Fprintln(writer, s[2]*s[2]+(s[0]+s[1])*(s[0]+s[1]))
	}
	writer.Flush()
}
