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

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		a := 0
		for ; (a+1)+(a+1)*(a+1) <= tmp; a++ {

		}
		fmt.Fprintln(writer, a)
	}
	writer.Flush()
}
