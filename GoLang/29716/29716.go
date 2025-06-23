package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var j, n int
	fmt.Fscanln(reader, &j, &n)

	var s string
	var score int
	scores := make([]int, n)

	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		score = 0
		for _, c := range s {
			if c == ' ' {
				score++
			} else if c >= 'A' && c <= 'Z' {
				score += 4
			} else if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				score += 2
			}
		}
		scores[i] = score
	}

	count := 0
	for i := 0; i < n; i++ {
		if scores[i] <= j {
			count++
		}
	}
	fmt.Println(count)
}
