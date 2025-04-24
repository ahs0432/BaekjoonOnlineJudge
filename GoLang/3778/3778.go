package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, ans int
	fmt.Fscanln(reader, &n)

	var a, b string
	check := make([][]int, 2)
	for i := 1; i <= n; i++ {
		ans = 0
		a, b = "", ""
		fmt.Fscanln(reader, &a)
		fmt.Fscanln(reader, &b)

		check[0] = make([]int, 26)
		check[1] = make([]int, 26)

		for _, c := range a {
			check[0][c-'a']++
		}

		for _, c := range b {
			check[1][c-'a']++
		}

		for j := 0; j < 26; j++ {
			if check[0][j] > check[1][j] {
				ans += check[0][j] - check[1][j]
			} else {
				ans += check[1][j] - check[0][j]
			}
		}
		fmt.Fprintf(writer, "Case #%d: %d\n", i, ans)
	}
	writer.Flush()
}
