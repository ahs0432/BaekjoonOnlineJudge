package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := 0, 0
	var a, b int
	for i := 0; i < 4; i++ {
		fmt.Fscanln(reader, &a, &b)
		n = n - a + b

		if m < n {
			m = n
		}
	}
	fmt.Println(m)
}
