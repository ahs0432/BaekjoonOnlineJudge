package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y int
	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &y)

	for ; x <= y; x += 60 {
		fmt.Fprintf(writer, "All positions change in year %d\n", x)
	}
	writer.Flush()
}
