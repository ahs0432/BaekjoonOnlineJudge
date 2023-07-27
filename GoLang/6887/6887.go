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

	ans := 1
	for i := 1; i*i <= n; i++ {
		ans = i
	}
	fmt.Printf("The largest square has side length %d.\n", ans)
}
