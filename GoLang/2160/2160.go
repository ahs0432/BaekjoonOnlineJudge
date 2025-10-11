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

	var table [5][50]string

	for l := 0; l < n; l++ {
		for i := 0; i < 5; i++ {
			fmt.Fscanln(reader, &table[i][l])
		}
	}

	var c, t1, t2 int
	ans := 987654321
	for l := 0; l < n; l++ {
		for b := l + 1; b < n; b++ {
			c = 0
			for i := 0; i < 5; i++ {
				for j := 0; j < 7; j++ {
					if table[i][l][j] != table[i][b][j] {
						c++
					}
					if c > ans {
						break
					}
				}
			}

			if c < ans {
				ans = c
				t1 = l + 1
				t2 = b + 1
			}
		}
	}

	fmt.Println(t1, t2)
}
