package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := make([]int, 5)
	fmt.Fscanln(reader, &n[0], &n[1], &n[2], &n[3], &n[4])

	var count int
	min := min(n[0], n[1], n[2], n[3], n[4])
	for {
		count = 0
		for _, i := range n {
			if min%i == 0 {
				count += 1
			}
		}
		if count > 2 {
			break
		}
		min++
	}
	fmt.Println(min)
}
