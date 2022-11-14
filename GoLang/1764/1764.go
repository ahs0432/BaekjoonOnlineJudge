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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	d := map[string]bool{}

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		d[tmp] = true
	}

	answer := make([]string, 0)

	for i := 0; i < m; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		if d[tmp] {
			answer = append(answer, tmp)
		}
	}

	sort.Slice(answer, func(i int, j int) bool {
		return answer[i] < answer[j]
	})

	fmt.Fprintln(writer, len(answer))
	for _, s := range answer {
		fmt.Fprintln(writer, s)
	}
	writer.Flush()
}
