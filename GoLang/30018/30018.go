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

	t := make([][]int, 2)
	ans := 0
	for i := 0; i < 2; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &t[i][j])

			if i == 1 && t[0][j]-t[i][j] > 0 {
				ans += t[0][j] - t[i][j]
			}
		}
	}
	fmt.Println(ans)
}
