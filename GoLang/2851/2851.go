package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	mus := make([]int, 10)
	score := 0
	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &mus[i])
	}
	for i := 0; i < 10; i++ {
		score += mus[i]
		if score >= 100 {
			if score-100 > 100-(score-mus[i]) {
				score -= mus[i]
			}
			break
		}
	}
	fmt.Println(score)
}
