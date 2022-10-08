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

	top := 0
	count := 0
	for i := 0; i < 9; i++ {
		var num1 int
		fmt.Fscanln(reader, &num1)
		if top < num1 {
			top = num1
			count = i + 1
		}
	}

	fmt.Fprintln(writer, top)
	fmt.Fprintln(writer, count)
}
