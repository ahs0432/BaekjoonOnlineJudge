package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var p int
	fmt.Fscanln(reader, &p)

	var n int
	var d, a, b, f float64
	for i := 0; i < p; i++ {
		fmt.Fscanln(reader, &n, &d, &a, &b, &f)
		fmt.Fprintf(writer, "%d %.6f\n", n, d/(a+b)*f)
	}
	writer.Flush()
}
