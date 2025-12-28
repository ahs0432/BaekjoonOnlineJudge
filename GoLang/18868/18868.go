package main

import (
	"bufio"
	"fmt"
	"os"
)

func pCompare(p1, p2 []int) bool {
	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			if p1[i] == p1[j] {
				if p2[i] != p2[j] {
					return false
				}
			} else if p1[i] > p1[j] {
				if p2[i] <= p2[j] {
					return false
				}
			} else if p1[i] < p1[j] {
				if p2[i] >= p2[j] {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscanln(reader, &m, &n)

	p := make([][]int, m)
	for i := 0; i < m; i++ {
		p[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &p[i][j])
			} else {
				fmt.Fscan(reader, &p[i][j])
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if pCompare(p[i], p[j]) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
