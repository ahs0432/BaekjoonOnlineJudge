package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	c, d := make([]int, 3), make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &c[i], &d[i])
	}

	var h int
	fmt.Fscanln(reader, &h)

	ans := 0

	for {
		for i := 0; i < 3; i++ {
			if ans%c[i] == 0 {
				h -= d[i]
			}
		}

		if h <= 0 {
			fmt.Println(ans)
			break
		} else {
			ans++
		}
	}
}
