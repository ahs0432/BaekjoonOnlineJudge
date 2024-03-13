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

	var tmp int
	now := 0
	max := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp != 0 {
			now++
			if now > max {
				max = now
			}
		} else {
			now = 0
		}
	}
	fmt.Println(max)
}
