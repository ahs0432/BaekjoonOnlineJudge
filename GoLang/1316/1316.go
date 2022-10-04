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

	answer := 0

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanln(reader, &str)
		list := map[rune]bool{}

		check := true

		for j, c := range str {
			if j == 0 {
				list[c] = true
				continue
			}

			if _, exists := list[c]; !exists {
				list[c] = true
			} else {
				if rune(str[j-1]) != c {
					check = false
					break
				}
			}
		}

		if check {
			answer++
		}
	}

	fmt.Println(answer)
}
