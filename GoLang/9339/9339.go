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

	var k, n, tmp int
	var d, h, m int
	var rt, rc, ans, r int
	var table map[int]bool
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &k)

		table = map[int]bool{}
		for j := 0; j < k; j++ {
			if j == k-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			table[tmp] = true
		}

		fmt.Fscanln(reader, &n)

		rt = 1441
		rc = 0
		ans = 0

		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &d, &h, &m)
			if h == -1 && m == -1 {
				continue
			}

			r = (h * 60) + m
			if table[d] && r <= 360 {
				rc++

				if r < rt {
					ans = d
					rt = r
				}
			}
		}
		fmt.Fprintln(writer, ans, rc)
	}
	writer.Flush()
}
