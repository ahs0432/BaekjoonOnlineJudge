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

	r := 1
	t := make([]int, n)
	for i := 0; i < n; i++ {
		if n == i {
			fmt.Fscanln(reader, &t[i])
		} else {
			fmt.Fscan(reader, &t[i])
		}

		if i != 0 && t[i-1] <= t[i] {
			r++
		}
	}
	fmt.Println(r)
}
