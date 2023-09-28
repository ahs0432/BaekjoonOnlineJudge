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

	for i := 0; i < n; i++ {
		y, k := 0, 0
		var yt, kt int

		for j := 0; j < 9; j++ {
			fmt.Fscanln(reader, &yt, &kt)
			y += yt
			k += kt
		}

		if y > k {
			fmt.Fprintln(writer, "Yonsei")
		} else if y < k {
			fmt.Fprintln(writer, "Korea")
		} else {
			fmt.Fprintln(writer, "Draw")
		}
	}

	writer.Flush()
}
