package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}

		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
	}

	fmt.Println(dp[n])
}
