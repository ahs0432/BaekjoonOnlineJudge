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

	if n == 0 {
		fmt.Println(0)
		return
	} else if n == 1 {
		fmt.Println(1)
		return
	}

	fib := make([]int, n+1)
	fib[0], fib[1], fib[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println(fib[n])
}
