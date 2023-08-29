package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n string
	fmt.Fscanln(reader, &n)

	for i := 0; i <= len(n); i += 10 {
		if i+10 >= len(n) {
			fmt.Fprintln(writer, n[i:len(n)])
		} else {
			fmt.Fprintln(writer, n[i:i+10])
		}
	}
	writer.Flush()
}
