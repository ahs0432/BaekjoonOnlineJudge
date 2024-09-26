package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, mm, t, r int
	fmt.Fscanln(reader, &n, &m, &mm, &t, &r)

	if m+t > mm {
		fmt.Println(-1)
	} else {
		now := m
		time := 0
		for {
			if now+t <= mm {
				n--
				now = now + t
			} else {
				now = now - r
				if now < m {
					now = m
				}
			}
			time++

			if n == 0 {
				break
			}
		}

		fmt.Println(time)
	}
}
