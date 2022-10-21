package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]int, n)

	for i := 0; i < n; i++ {
		table[i] = i + 1
	}

	answer := make([]int, 0)

	now := -1
	for len(table) != 0 {
		now += k

		if now >= len(table) {
			now = now % len(table)
		}

		answer = append(answer, table[now])
		table = append(table[:now], table[(now+1):]...)
		now--
	}

	fmt.Print("<")
	for n, i := range answer {
		fmt.Print(i)
		if n != len(answer)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println(">")
		}
	}
}
