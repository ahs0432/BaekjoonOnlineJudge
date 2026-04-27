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
	fmt.Fscan(reader, &n)

	odd := 0
	even := 0
	now := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var val int
			fmt.Fscan(reader, &val)

			if val == 1 {
				if (i+j)%2 == 1 {
					odd++
				} else {
					even++
				}
			} else if val == 2 {
				now = (i + j) % 2
			}
		}
	}

	if (even == 0 && now == 0) || (odd == 0 && now == 1) {
		fmt.Fprintln(writer, "Lena")
	} else {
		fmt.Fprintln(writer, "Kiriya")
	}
	writer.Flush()
}
