package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibonacci(n int) int {
	f := make([]int, n+1)

	f[0] = 0
	f[1] = 1
	f[2] = 1

	c := 0
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
		c++
	}

	return c
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return (fib(n-1) + fib(n-2))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	fmt.Println(fib(n), fibonacci(n))
}
