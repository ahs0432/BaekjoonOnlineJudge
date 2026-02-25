package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	var q int
	fmt.Fscanln(reader, &q)

	var t, l, r, k, count int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t)

		if t == 1 {
			fmt.Fscanln(reader, &l, &r, &k)
			count = 0
			for j := l - 1; j < r; j++ {
				if a[j] == k {
					count++
				}
			}
			fmt.Fprintln(writer, count)
		} else if t == 2 {
			fmt.Fscanln(reader, &l, &r)
			for j := l - 1; j < r; j++ {
				a[j] = -1
			}
		}
	}
	writer.Flush()
}
