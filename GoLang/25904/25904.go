package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscanln(reader, &n, &x)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	for i := 0; ; i++ {
		if table[i%n] < x+i {
			fmt.Println(i%n + 1)
			break
		}
	}
}
