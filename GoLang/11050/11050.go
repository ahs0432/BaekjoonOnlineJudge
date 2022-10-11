package main

import (
	"bufio"
	"fmt"
	"os"
)

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	nk := factorial(n - k)
	n = factorial(n)
	k = factorial(k)

	answer := n/(k*nk) + n%(k*nk)
	fmt.Println(answer)
}
