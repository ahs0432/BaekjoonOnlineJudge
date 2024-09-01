package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	table := make([]int, 101)

	var s, e int
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &s, &e)
		for ; s < e; s++ {
			table[s]++
		}
	}

	ans := 0
	for i := 0; i < 101; i++ {
		if table[i] == 1 {
			ans += table[i] * a
		} else if table[i] == 2 {
			ans += table[i] * b
		} else if table[i] == 3 {
			ans += table[i] * c
		}
	}
	fmt.Println(ans)
}
