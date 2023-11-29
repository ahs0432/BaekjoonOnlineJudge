package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	h, a := 0, 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if h < tmp {
			h = tmp
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &tmp)
		if a < tmp {
			a = tmp
		}
	}

	fmt.Println(h + a)
}
