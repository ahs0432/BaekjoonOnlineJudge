package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	arr := make([]int, 1004)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var box, a, b, c, d int
	var s1, s2 int64
	for i := 0; i < q; i++ {
		s1, s2 = 0, 0
		fmt.Fscan(reader, &box)

		if box == 1 {
			fmt.Fscanln(reader, &a, &b)
			for j := a; j <= b; j++ {
				s1 += (int64(arr[j]))
			}
			fmt.Fprintln(writer, s1)
			arr[a], arr[b] = arr[b], arr[a]
		} else if box == 2 {
			fmt.Fscanln(reader, &a, &b, &c, &d)
			for j := a; j <= b; j++ {
				s1 += (int64(arr[j]))
			}
			for j := c; j <= d; j++ {
				s2 += (int64(arr[j]))
			}
			fmt.Fprintln(writer, s1-s2)
		}
	}
	writer.Flush()
}
