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

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	ans := 1
	now := 2
	table[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if table[i] > now {
			table[i] = now
		} else {
			now = table[i]
		}
		ans += table[i]
		now++
	}

	fmt.Println(ans)
}
