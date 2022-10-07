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
		j := 0
		for ; j < num1-i; j++ {
			fmt.Fprint(writer, " ")
		}
		for ; j < num1; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer, "")
	}
}
