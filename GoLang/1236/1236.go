package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var y, x int
	fmt.Fscanln(reader, &y, &x)

	table := make([]string, y)
	for i := 0; i < y; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	a, b := 0, 0
	for i := 0; i < y; i++ {
		if !strings.Contains(table[i], "X") {
			a++
		}
	}

	check := true
	for i := 0; i < x; i++ {
		check = true
		for j := 0; j < y; j++ {
			if table[j][i] == 'X' {
				check = false
				break
			}
		}

		if check {
			b++
		}
	}

	fmt.Println(max(a, b))

}
