package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n float64
	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		if n <= 1000000 {
			fmt.Fprintln(writer, int(n))
		} else if n <= 5000000 {
			fmt.Fprintln(writer, int(n*0.9))
		} else {
			fmt.Fprintln(writer, int(n*0.8))
		}
	}
	writer.Flush()
}
