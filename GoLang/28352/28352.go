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

	a := 1

	for i := 2; i <= n; i++ {
		a *= i
	}

	ans := 1
	for i := 2; (i * 7 * 24 * 60 * 60) <= a; i++ {
		ans = i
	}
	fmt.Println(ans)
}
