package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	var l, i, sum int
	t := make([]bool, n)
	for a := 0; a < q; a++ {
		fmt.Fscanln(reader, &l, &i)

		for l = l - 1; l < n; l += i {
			if !t[l] {
				t[l] = true
				sum++
			}
		}
	}
	fmt.Println(n - sum)
}
