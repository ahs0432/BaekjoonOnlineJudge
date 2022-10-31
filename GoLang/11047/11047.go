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

	coins := make([]int, n)

	able := n - 1

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &coins[i])
		if coins[i] > m {
			able = i - 1
		}
	}

	answer := 0
	for ; able >= 0; able-- {
		answer += m / coins[able]
		m %= coins[able]
	}

	fmt.Println(answer)
}
