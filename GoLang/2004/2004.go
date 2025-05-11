package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(a, b int) int {
	res := 0
	tmp := b

	for tmp <= a {
		res += a / tmp
		tmp *= b
	}

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	twoCount := count(n, 2) - count(m, 2) - count(n-m, 2)
	fiveCount := count(n, 5) - count(m, 5) - count(n-m, 5)

	fmt.Println(min(twoCount, fiveCount))
}
