package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ans := 0
	table := make([]string, 8)
	for i := 0; i < 8; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 && table[i][j] == 'F' {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
