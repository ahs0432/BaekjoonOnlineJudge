package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var num1 int
	fmt.Fscanln(reader, &num1)

	for i := 1; i <= num1; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer, "")
	}
}
