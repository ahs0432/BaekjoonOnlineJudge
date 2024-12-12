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

	boxes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boxes[i])
	}

	books := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &books[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if boxes[j] < books[i] {
				continue
			} else {
				boxes[j] -= books[i]
				break
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += boxes[i]
	}
	fmt.Println(sum)
}
