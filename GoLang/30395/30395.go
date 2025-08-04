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

	stf := 1
	var tmp, max, now int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp == 0 {
			if stf >= 1 {
				stf = -1
			} else {
				now = 0
			}
		} else {
			now++

			if now > max {
				max = now
			}
		}

		stf++
	}
	fmt.Println(max)
}
