package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m, &k)

	var f, d, sum int
	min := 2147483647
	ans := 0

	for i := 1; i <= k; i++ {
		fmt.Fscanln(reader, &f, &d)
		sum = f - 1 + m - d

		if min > sum {
			ans = i
			min = sum
		}
	}
	fmt.Println(ans)
}
