package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for reader.Scan() {
		if s := reader.Text(); s == "" {
			break
		} else {
			fmt.Fprintln(writer, s)
		}
	}
	writer.Flush()
}
