package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var k int
	fmt.Fscanln(reader, &k)

	for i := 0; i < k; i++ {
		var n, s, d int
		fmt.Fscanln(reader, &n, &s, &d)

		du := 0
		for j := 0; j < n; j++ {
			var di, vi int
			fmt.Fscanln(reader, &di, &vi)

			if s*d >= di {
				du += vi
			}
		}

		fmt.Fprintf(writer, "Data Set %d:\n", i+1)
		fmt.Fprint(writer, du, "\n\n")
	}

	writer.Flush()
}
