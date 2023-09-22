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

	y, m := 0, 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		y += (tmp/30 + 1) * 10
		m += (tmp/60 + 1) * 15
	}

	if y == m {
		fmt.Println("Y M", y)
	} else if y > m {
		fmt.Println("M", m)
	} else {
		fmt.Println("Y", y)
	}
}
