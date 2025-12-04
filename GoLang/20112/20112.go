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
	table := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	check := true
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if table[i][j] != table[j][i] {
				check = false
			}
		}
	}

	if check {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
