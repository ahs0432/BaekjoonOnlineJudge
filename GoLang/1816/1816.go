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

	var s int64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)

		for j := int64(2); j < 1000001; j++ {
			if s%j == 0 {
				fmt.Fprintln(writer, "NO")
				break
			}

			if j == 1000000 {
				fmt.Fprintln(writer, "YES")
			}
		}
	}
	writer.Flush()
}
