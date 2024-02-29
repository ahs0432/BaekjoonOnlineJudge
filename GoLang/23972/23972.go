package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, n int
	fmt.Fscanln(reader, &k, &n)

	if n == 1 {
		fmt.Println(-1)
	} else {
		ans := (k * n) / (n - 1)
		if (k*n)%(n-1) != 0 {
			ans++
		}
		fmt.Println(ans)
	}
}
