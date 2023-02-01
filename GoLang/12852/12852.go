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
	dpPass := make([][]int, n+1)
	dpPass[1] = []int{1}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		dpPass[i] = make([]int, len(dpPass[i-1]))
		copy(dpPass[i], dpPass[i-1])

		if i%3 == 0 && dp[i] > dp[i/3]+1 {
			dp[i] = dp[i/3] + 1
			dpPass[i] = make([]int, len(dpPass[i/3]))
			copy(dpPass[i], dpPass[i/3])
		}

		if i%2 == 0 && dp[i] > dp[i/2]+1 {
			dp[i] = dp[i/2] + 1
			dpPass[i] = make([]int, len(dpPass[i/2]))
			copy(dpPass[i], dpPass[i/2])
		}

		dpPass[i] = append(dpPass[i], i)
	}

	fmt.Println(dp[n])
	for i := len(dpPass[n]) - 1; i >= 0; i-- {
		fmt.Print(dpPass[n][i], " ")
	}
}
