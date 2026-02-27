package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	var tmp, a []int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		tmp = make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = j + 1
		}

		a = make([]int, n)
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &a[j])
			} else {
				fmt.Fscan(reader, &a[j])
			}
		}

		for j := 0; j < n; j++ {
			if tmp[j] == a[j] {
				tmp[j]++
			}

			for k := j + 1; k < n; k++ {
				if tmp[k-1] >= tmp[k] {
					tmp[k] += (tmp[k-1] - tmp[k]) + 1
				}
			}
		}
		fmt.Fprintln(writer, tmp[len(tmp)-1])
	}
	writer.Flush()
}
