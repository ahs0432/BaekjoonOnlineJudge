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

	for i := 0; i < num1; i++ {
		var num2, num3 int
		fmt.Fscanln(reader, &num2, &num3)
		fmt.Fprintln(writer, num2+num3)
	}
}
