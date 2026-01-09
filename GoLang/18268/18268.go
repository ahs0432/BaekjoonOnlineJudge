package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, n, tmp int
	fmt.Fscanln(reader, &k, &n)

	var rank [11][21]int
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			rank[i][tmp] = j
		}
	}

	ans := 0
	var check bool
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}

			check = true
			for l := 0; l < k; l++ {
				if rank[l][i] > rank[l][j] {
					check = false
				}
			}
			if check {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
