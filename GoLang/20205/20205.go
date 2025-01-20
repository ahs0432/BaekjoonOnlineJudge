package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	now := ""
	var tmp string
	for i := 0; i < n; i++ {
		now = ""
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			for a := 0; a < k; a++ {
				now += tmp + " "
			}
		}

		for a := 0; a < k; a++ {
			fmt.Fprintln(writer, now)
		}
	}

	writer.Flush()
}
