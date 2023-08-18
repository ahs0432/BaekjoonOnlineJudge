package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	s := make([]string, n)
	a := 0

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s[i])
	}
	fmt.Fscanln(reader)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		for j := 0; j < m; j++ {
			if s[i][j] == tmp[j] {
				a++
			}
		}
	}

	fmt.Println(a)
}
