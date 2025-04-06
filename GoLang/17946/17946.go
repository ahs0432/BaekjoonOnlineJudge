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

	var tmp int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, 1)
	}
	writer.Flush()
}
