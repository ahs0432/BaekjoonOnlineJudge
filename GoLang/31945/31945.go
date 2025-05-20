package main

import (
	"bufio"
	"fmt"
	"os"
)

var p [][]int = [][]int{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}, {0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}}

func checkP(tmp []int) bool {
	sum := []int{0, 0, 0}

	for _, t := range tmp {
		point := p[t]
		for i := 0; i < 3; i++ {
			sum[i] += point[i]
		}
	}

	c0, c2, c4 := 0, 0, 0

	for _, a := range sum {
		if a == 0 {
			c0++
		} else if a == 2 {
			c2++
		} else if a == 4 {
			c4++
		}
	}

	return c2 == 2 && (c0 == 1 || c4 == 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	tmp := make([]int, 4)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &tmp[0], &tmp[1], &tmp[2], &tmp[3])
		if checkP(tmp) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
