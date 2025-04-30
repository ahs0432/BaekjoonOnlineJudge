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

	var b, p float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &b, &p)
		fmt.Fprintf(writer, "%.4f %.4f %.4f\n", 60/p*(b-1), 60/p*b, 60/p*(b+1))
	}
	writer.Flush()
}
