package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	sum := 0
	max := 0

	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &a, &b)

		sum = sum - a + b
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
}
