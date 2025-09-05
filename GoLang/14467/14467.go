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

	cows := make([]int, n+1)
	nows := make([]int, n+1)

	for i := 0; i <= n; i++ {
		nows[i] = -1
	}

	var now, pos, sum int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &now, &pos)
		if nows[now] == -1 {
			nows[now] = pos
		} else if nows[now] != pos {
			sum++
			cows[now]++
			nows[now] = pos
		}
	}
	fmt.Println(sum)
}
