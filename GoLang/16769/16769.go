package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	c, m := make([]int, 3), make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &c[i], &m[i])
	}

	for i := 0; i < 100; i++ {
		if m[i%3]+m[(i+1)%3] > c[(i+1)%3] {
			m[i%3] = m[i%3] + m[(i+1)%3] - c[(i+1)%3]
			m[(i+1)%3] = c[(i+1)%3]
		} else {
			m[(i+1)%3] = m[i%3] + m[(i+1)%3]
			m[i%3] = 0
		}
	}

	for i := 0; i < 3; i++ {
		fmt.Println(m[i])
	}
}
