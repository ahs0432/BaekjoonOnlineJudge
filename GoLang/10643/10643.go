package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	pizza := make([]int, 8)
	for i := 0; i < 8; i++ {
		fmt.Fscanln(reader, &pizza[i])
	}

	max := 0
	now := 0
	for i := 0; i < 8; i++ {
		now = 0

		if i+4 < 8 {
			for j := i; j < i+4; j++ {
				now += pizza[j]
			}
		} else {
			for j := i; j < 8; j++ {
				now += pizza[j]
			}
			for j := 0; j < i+4-8; j++ {
				now += pizza[j]
			}
		}

		if now > max {
			max = now
		}
	}
	fmt.Println(max)
}
