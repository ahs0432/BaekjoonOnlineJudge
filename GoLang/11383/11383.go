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

	s1 := make([]string, n)
	s2 := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s1[i])
	}

	check := true
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s2[i])
		if check {
			if len(s1[i])*2 != len(s2[i]) {
				check = false
				continue
			}

			for j := 0; j < m; j++ {
				if s1[i][j] != s2[i][j*2] || s1[i][j] != s2[i][j*2+1] {
					check = false
					break
				}
			}
		}
	}

	if check {
		fmt.Println("Eyfa")
	} else {
		fmt.Println("Not Eyfa")
	}
}
