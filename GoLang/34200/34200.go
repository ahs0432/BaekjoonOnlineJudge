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

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	for i := 1; i < n; i++ {
		if x[i]-x[i-1] < 2 {
			fmt.Println(-1)
			return
		}
	}

	move := 0
	cur := 0

	for _, tmp := range x {
		for tmp-cur >= 3 {
			cur += 2
			move++
		}

		if tmp-cur == 2 {
			cur += 1
			move++
		}

		cur += 2
		move++
	}

	fmt.Println(move)
}
