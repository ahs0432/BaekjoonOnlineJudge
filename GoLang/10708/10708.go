package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := make([]int, n)
	b := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &b[j])
		}

		for j := 0; j < n; j++ {
			if b[j] == a[i] {
				ans[j]++
			} else {
				ans[a[i]-1]++
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, ans[i])
	}
	writer.Flush()
}
