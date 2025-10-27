package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int64
	fmt.Fscanln(reader, &n)

	var count int
	if n > 0 {
		count = 1
		k = 1
		for k < n {
			k *= 2
			count++
		}
	}
	fmt.Println(count)
}
