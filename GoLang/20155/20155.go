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

	ans := make([]int, m+1)

	max := -1
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		ans[tmp]++

		if ans[tmp] > max {
			max = ans[tmp]
		}
	}
	fmt.Println(max)
}
