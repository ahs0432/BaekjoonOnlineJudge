package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		reader.Scan()
		if s := reader.Text(); len(s) != 0 {
			fmt.Fprintln(writer, s)
		} else {
			break
		}
	}
	writer.Flush()
}
