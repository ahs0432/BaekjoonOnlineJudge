package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, n int
	fmt.Fscanln(reader, &l)
	fmt.Fscanln(reader, &n)

	var p, k int
	max, now := -1, -1
	cake := make([]int, 1000)
	count := make([]int, 1000)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &p, &k)
		for j := p; j <= k; j++ {
			if cake[j] == 0 {
				cake[j] = i
			}
		}

		if k-p+1 > max {
			max = k - p + 1
			now = i
		}
	}
	fmt.Println(now)

	for i := 1; i <= l; i++ {
		if cake[i] > 0 {
			count[cake[i]]++
		}
	}

	max, now = -1, -1
	for i := 1; i <= n; i++ {
		if count[i] > max {
			max = count[i]
			now = i
		}
	}
	fmt.Println(now)
}
