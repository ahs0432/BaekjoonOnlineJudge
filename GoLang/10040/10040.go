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

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
		fmt.Fscanln(reader, &a[i][0])
	}

	var tmp int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &tmp)

		for j := 0; j < n; j++ {
			if a[j][0] <= tmp {
				a[j][1]++
				break
			}
		}
	}

	var ans, max int
	for i := 0; i < n; i++ {
		if max < a[i][1] {
			ans = i + 1
			max = a[i][1]
		}
	}

	fmt.Println(ans)
}
