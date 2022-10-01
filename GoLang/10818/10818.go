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

	max := 0
	min := 0
	first := true
	for i := 0; i < num1; i++ {
		var num2 int
		fmt.Fscanf(reader, "%d ", &num2)
		if first {
			max = num2
			min = num2
			first = false
		} else {
			if max < num2 {
				max = num2
			} else if min > num2 {
				min = num2
			}
		}
	}

	fmt.Fprintf(writer, "%d %d", min, max)
}
