package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d, h, m int
	fmt.Fscanln(reader, &d, &h, &m)

	sum := 16511
	h += (d * 24)
	m += (h * 60)

	if m-sum < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(m - sum)
	}
}
