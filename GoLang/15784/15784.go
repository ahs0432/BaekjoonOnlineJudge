package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscanln(reader, &n, &x, &y)

	t := make([][]int, n)
	x--
	y--

	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &t[i][j])
		}
	}

	for i := 0; i < n; i++ {
		if t[i][y] > t[x][y] {
			fmt.Println("ANGRY")
			return
		}
		if t[x][i] > t[x][y] {
			fmt.Println("ANGRY")
			return
		}
	}

	fmt.Println("HAPPY")
}
