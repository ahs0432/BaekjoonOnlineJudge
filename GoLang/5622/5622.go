package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	list := map[rune]int{}

	for i := 1; i < 11; i++ {
		list[rune(i%10+48)] = (i + 2)
	}

	now := 65
	count := []int{3, 3, 3, 3, 3, 4, 3, 4}

	for n, c := range count {
		for i := 0; i < c; i++ {
			list[rune(now)] = (n + 3)
			now += 1
		}
	}

	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &s)

	answer := 0

	for _, c := range s {
		answer += list[c]
	}

	fmt.Println(answer)
}
