package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	c := 0
	cnt := 0
	for i := 512; i >= 1; i /= 2 {
		cnt = 0
		if i <= a {
			cnt++
			a -= i
		}

		if i <= b {
			cnt++
			b -= i
		}

		if cnt == 1 {
			c += i
		}
	}

	fmt.Println(c)
}
