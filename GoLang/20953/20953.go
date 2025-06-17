package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, (a+b-1)*(a+b)*(a+b)/2)
	}
	writer.Flush()
}
