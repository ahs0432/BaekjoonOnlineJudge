package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x, s int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &x, &s)

	ans := "NO"
	weapon := make([][]int, n)
	for i := 0; i < n; i++ {
		weapon[i] = make([]int, 2)
		fmt.Fscanln(reader, &weapon[i][0], &weapon[i][1])

		if weapon[i][0] <= x && weapon[i][1] > s {
			ans = "YES"
		}
	}

	fmt.Println(ans)
}
