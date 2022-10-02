package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	var num1, num2 int
	fmt.Fscanln(reader, &num1, &num2)
	defer writer.Flush()

	for i := 0; i < num1; i++ {
		var num3 int
		fmt.Fscanf(reader, "%d ", &num3)

		if num3 < num2 {
			fmt.Fprintf(writer, "%d ", num3)
		}

	}
}
