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

	counts := make(map[int]int)

	var k, tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &k)

		for j := 0; j < k; j++ {
			if j == k-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}
			counts[tmp]++
		}
	}

	ans := 0
	for _, count := range counts {
		if count >= m {
			ans++
		}
	}
	fmt.Println(ans)
}
