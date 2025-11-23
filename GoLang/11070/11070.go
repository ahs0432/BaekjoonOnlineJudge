package main

import (
	"bufio"
	"fmt"
	"os"
)

func sol(s, a int64) int64 {
	if s == 0 && a == 0 {
		return 0
	} else {
		return (s * s) * 1000 / (s*s + a*a)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, m int
	var a, b, p, q int64
	var min, max int64

	var S, A, W []int64
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &m)

		S = make([]int64, 1001)
		A = make([]int64, 1001)
		W = make([]int64, 1001)
		for i := 0; i < m; i++ {
			fmt.Fscanln(reader, &a, &b, &p, &q)
			S[a] += p
			A[a] += q
			S[b] += q
			A[b] += p
		}

		min, max = 0, 0
		for i := 1; i <= n; i++ {
			W[i] = sol(S[i], A[i])
			if i == 1 {
				max = W[i]
				min = W[i]
			} else {
				if max < W[i] {
					max = W[i]
				}
				if min > W[i] {
					min = W[i]
				}
			}
		}
		fmt.Fprintf(writer, "%d\n%d\n", max, min)
	}
	writer.Flush()
}
