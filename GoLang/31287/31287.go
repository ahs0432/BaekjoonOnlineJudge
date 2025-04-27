package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	if k > n/2 {
		k = n / 2
	}

	var s string
	fmt.Fscanln(reader, &s)

	nowX, nowY := 0, 0
	for i := 0; i < k; i++ {
		for _, c := range s {
			if c == 'L' {
				nowX--
			} else if c == 'R' {
				nowX++
			} else if c == 'U' {
				nowY++
			} else {
				nowY--
			}

			if nowX == 0 && nowY == 0 {
				fmt.Println("YES")
				return
			}
		}

	}
	fmt.Println("NO")
}
