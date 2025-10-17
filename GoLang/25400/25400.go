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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	s := 1
	count := 0

	for _, i := range a {
		if i == s {
			s++
		} else {
			count++
		}
	}
	fmt.Println(count)
}
