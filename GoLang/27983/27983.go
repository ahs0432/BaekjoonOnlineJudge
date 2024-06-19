package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	x, l := make([]int, n), make([]int, n)
	c := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if abs(x[i]-x[j]) > l[i]+l[j] || c[i] == c[j] {
				continue
			}
			fmt.Println("YES")
			fmt.Println(i+1, j+1)
			return
		}
	}
	fmt.Println("NO")
}
