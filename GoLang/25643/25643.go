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
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s[i])
	}

	var check bool
	for i := 0; i < n-1; i++ {
		check = false
		for j := 1; j < m+1; j++ {
			if s[i][m-j:] == s[i+1][:j] {
				check = true
				break
			}

			if s[i][:j] == s[i+1][m-j:] {
				check = true
				break
			}
		}

		if !check {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)
}
