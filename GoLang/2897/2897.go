package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, c int
	fmt.Fscanln(reader, &r, &c)

	table := make([]string, r)
	for i := 0; i < r; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	var count int
	ans := make([]int, 5)
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			count = 0
			if table[i-1][j-1] == '#' {
				continue
			}
			if table[i-1][j-1] == 'X' {
				count++
			}
			if table[i][j-1] == '#' {
				continue
			}
			if table[i][j-1] == 'X' {
				count++
			}
			if table[i-1][j] == '#' {
				continue
			}
			if table[i-1][j] == 'X' {
				count++
			}
			if table[i][j] == '#' {
				continue
			}
			if table[i][j] == 'X' {
				count++
			}
			ans[count]++
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(ans[i])
	}
}
