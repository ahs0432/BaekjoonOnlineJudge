package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	table := make([]int, 26)

	for _, c := range s {
		table[c-97]++
	}

	for _, i := range table {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
