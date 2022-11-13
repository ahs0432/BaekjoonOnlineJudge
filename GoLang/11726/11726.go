package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	if n > 1 {
		dp[2] = 2

		for i := 3; i <= n; i++ {
			dp[i] = (dp[i-1] + dp[i-2]) % 10007
		}
	}

	fmt.Println(dp[n])
}
