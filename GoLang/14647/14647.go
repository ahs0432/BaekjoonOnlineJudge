package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var sum, max int
	var tmp string
	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			for _, c := range tmp {
				if c == '9' {
					sum++
					row[i]++
					col[j]++
				}
			}

			if row[i] > max {
				max = row[i]
			}
			if col[j] > max {
				max = col[j]
			}
		}
	}

	fmt.Println(sum - max)
}
