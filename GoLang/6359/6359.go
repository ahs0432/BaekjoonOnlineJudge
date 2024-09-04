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

	var n, ans int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		ans = 0
		for j := 1; j < n; j++ {
			if j*j <= n {
				ans++
			} else {
				break
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
