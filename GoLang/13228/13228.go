package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var x1, y1, f1, x2, y2, f2 int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &x1, &y1, &f1, &x2, &y2, &f2)
		fmt.Fprintln(writer, f1+f2+abs(x2-x1)+abs(y2-y1))
	}
	writer.Flush()
}
