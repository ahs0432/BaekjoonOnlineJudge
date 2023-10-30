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

	var tmp float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintf(writer, "$%.2f\n", tmp*0.8)
	}
	writer.Flush()
}
