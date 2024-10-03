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

	count := 0
	for i := 1; i <= n; i++ {
		for j := i; i*j <= n; j++ {
			count++
		}
	}
	fmt.Println(count)
}
