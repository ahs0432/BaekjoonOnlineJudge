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
	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}
		ans = 0

		for i := 1; i <= n; i++ {
			ans += (n - i + 1) * (n - i + 1)
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
