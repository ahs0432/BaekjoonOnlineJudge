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

	table := make([]int, n+1)
	for i := 1; i <= n; i++ {
		table[i] = table[i-1] + i
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += table[i]
	}
	for i := n - 1; i > 0; i -= 2 {
		ans += table[i]
	}

	fmt.Println(ans)
}
