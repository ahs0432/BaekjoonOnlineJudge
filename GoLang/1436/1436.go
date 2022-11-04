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

	count := 1
	now := 666
	for count != n {
		now += 1
		t := 0
		tmp := now

		for tmp != 0 {
			if tmp%10 == 6 {
				if t == 2 {
					count++
					break
				}
				t++
			} else {
				t = 0
			}

			tmp /= 10
		}
	}

	fmt.Println(now)
}
