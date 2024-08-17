package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	table := make([]int, 9)

	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &table[i])
		sum += table[i]
	}

	breakPoint := false
	sum -= 100
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if table[i]+table[j]-sum == 0 {
				table = append(table[:i], table[i+1:]...)
				table = append(table[:j-1], table[j:]...)
				breakPoint = true
				break
			}
		}
		if breakPoint {
			break
		}
	}

	for i := 0; i < 7; i++ {
		fmt.Println(table[i])
	}
}
