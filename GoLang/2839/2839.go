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

	answer := n / 5
	tmp := n % 5

	if tmp != 0 {
		for answer >= 0 {
			if tmp%3 == 0 {
				answer += (tmp / 3)
				break
			}

			answer--
			tmp += 5
		}
	}

	fmt.Println(answer)
}
