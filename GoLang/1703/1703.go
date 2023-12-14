package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, ans, a, b int

	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}

		ans = 1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a, &b)
			ans *= a
			ans -= b
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
