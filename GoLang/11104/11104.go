package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var ans, now int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)

		ans = 0
		now = 1
		for j := len(s) - 1; j >= 0; j-- {
			ans += (int(s[j]) - 48) * now
			now *= 2
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
