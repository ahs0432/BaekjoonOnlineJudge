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

	pi := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &pi[i])
	}

	var max, tmp int
	for i := 0; i < n-1; i++ {
		if pi[i] < pi[i+1] {
			tmp += (pi[i+1] - pi[i])
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}

	if tmp > max {
		max = tmp
	}
	fmt.Println(max)
}
