package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, ans int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		ans = 1
		for j := 1; j <= n; j++ {
			ans *= j
			ans %= 100000
			for ans%10 == 0 {
				ans /= 10
			}
		}
		fmt.Fprintln(writer, ans%10)
	}
	writer.Flush()
}
