package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	for {
		_, err := fmt.Fscanln(reader, &n)
		if err == io.EOF {
			break
		}
		if n%6 == 0 {
			fmt.Fprintln(writer, "Y")
		} else {
			fmt.Fprintln(writer, "N")
		}
	}
	writer.Flush()
}
