package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]byte, n)
	fmt.Fscanln(reader, &table)

	var k int
	fmt.Fscanln(reader, &k)

	for i := 0; i < k; i++ {
		if len(table) == 0 {
			break
		}

		if table[len(table)-1] != '0' {
			fmt.Println("NO")
			return
		}
		table = table[:len(table)-1]
	}
	fmt.Println("YES")
}
