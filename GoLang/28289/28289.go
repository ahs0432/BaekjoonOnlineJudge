package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p int
	fmt.Fscanln(reader, &p)

	table := make([]int, 4)
	var g, c, n int
	for i := 0; i < p; i++ {
		fmt.Fscanln(reader, &g, &c, &n)

		if g == 1 {
			table[3]++
		} else if c <= 2 {
			table[0]++
		} else if c == 3 {
			table[1]++
		} else {
			table[2]++
		}
	}

	fmt.Println(table[0])
	fmt.Println(table[1])
	fmt.Println(table[2])
	fmt.Println(table[3])
}
