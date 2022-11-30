package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	floors := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &floors[i])
	}

	if t == 1 {
		fmt.Println(floors[0])
	} else if t == 2 {
		fmt.Println(floors[0] + floors[1])
	} else {
		dp := make([]int, t+1)
		dp[1] = floors[0]
		dp[2] = floors[0] + floors[1]

		for i := 3; i <= t; i++ {
			dp[i] = max(dp[i-3]+floors[i-2]+floors[i-1], dp[i-2]+floors[i-1])
		}

		fmt.Println(dp[t])
	}
}
