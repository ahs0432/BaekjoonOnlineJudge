package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	max := 0
	var n int
	table := make([]int, 1001)
	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &n)
		sum += n
		table[n]++

		if table[max] < table[n] {
			max = n
		}
	}

	fmt.Println(sum / 10)
	fmt.Println(max)
}
