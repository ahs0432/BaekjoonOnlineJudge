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
	for {
		fmt.Fscanln(reader, &n)

		if n == -1 {
			break
		}

		s := make([]int, n)
		t := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanln(reader, &s[i], &t[i])
		}

		ans := s[0] * t[0]
		for i := 1; i < n; i++ {
			ans += s[i] * (t[i] - t[i-1])
		}
		fmt.Fprintln(writer, ans, "miles")
	}
	writer.Flush()
}
