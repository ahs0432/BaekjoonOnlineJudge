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

	var t int
	fmt.Fscanln(reader, &t)

	n := make([]int, 5)
	var sum int
	for i := 0; i < t; i++ {
		for j := 0; j < 5; j++ {
			fmt.Fscan(reader, &n[j])
		}

		sort.Slice(n, func(k, l int) bool {
			return n[k] < n[l]
		})

		if n[3]-n[1] >= 4 {
			fmt.Fprintln(writer, "KIN")
			continue
		}

		sum = 0
		for j := 1; j < 4; j++ {
			sum += n[j]
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
