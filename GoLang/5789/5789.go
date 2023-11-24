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

	var tmp string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp[len(tmp)/2] == tmp[len(tmp)/2-1] {
			fmt.Fprintln(writer, "Do-it")
		} else {
			fmt.Fprintln(writer, "Do-it-Not")
		}
	}
	writer.Flush()
}
