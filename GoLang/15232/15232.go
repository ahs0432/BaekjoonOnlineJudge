package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var r, c int
	fmt.Fscanln(reader, &r)
	fmt.Fscanln(reader, &c)

	a := ""
	for i := 0; i < c; i++ {
		a += "*"
	}

	for i := 0; i < r; i++ {
		fmt.Fprintln(writer, a)
	}
	writer.Flush()
}
