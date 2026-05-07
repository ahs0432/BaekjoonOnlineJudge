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

	var k, ans int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &k)

		ans = 1
		for j := 1; j < k; j++ {
			ans = ans*2 + 1
		}
		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
