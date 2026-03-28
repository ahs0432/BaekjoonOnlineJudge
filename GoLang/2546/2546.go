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

	var n, m, ans int
	var sumC, sumE int64
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n, &m)

		sumC, sumE = 0, 0
		cTable := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cTable[j])
			sumC += int64(cTable[j])
		}

		eTable := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &eTable[j])
			sumE += int64(eTable[j])
		}

		ans = 0
		for i := 0; i < n; i++ {
			if int64(cTable[i])*int64(n) < sumC && int64(cTable[i])*int64(m) > sumE {
				ans++
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
