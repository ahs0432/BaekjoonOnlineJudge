package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	table := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscanln(reader, &a, &b)
		table[i] = b - a
	}

	fmt.Println(table[3] + table[2] + table[1])
	fmt.Println(table[3] + table[2])
	fmt.Println(table[3])
}
