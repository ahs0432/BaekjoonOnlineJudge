package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	var table []int
	tmp, count := 1, 1

	for i := 0; i < b; i++ {
		table = append(table, tmp)
		if tmp == count {
			count = 0
			tmp++
		}
		count++
	}

	sum := 0
	for i := a - 1; i < b; i++ {
		sum += table[i]
	}

	fmt.Println(sum)
}
