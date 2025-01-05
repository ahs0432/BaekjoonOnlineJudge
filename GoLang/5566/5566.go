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

	board := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &board[i])
	}

	now := 0
	var tmp int
	finish := -1
	for i := 1; i <= m; i++ {
		fmt.Fscanln(reader, &tmp)

		if finish != -1 {
			continue
		}

		now += tmp
		if now >= n {
			now = n - 1
		}

		now += board[now]
		if now >= n-1 {
			finish = i
		}
	}
	fmt.Println(finish)
}
