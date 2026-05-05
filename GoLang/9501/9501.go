package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, d, ans, vi, fi, ci int
	for i := 0; i < t; i++ {
		ans = 0
		fmt.Fscanln(reader, &n, &d)

		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &vi, &fi, &ci)
			if vi*fi/ci >= d {
				ans++
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
