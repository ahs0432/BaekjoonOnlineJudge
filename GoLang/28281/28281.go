package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, now int
	fmt.Fscanln(reader, &n, &m)
	fmt.Fscan(reader, &now)

	min := 2000
	for i := 1; i < n; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)

		if min > tmp+now {
			min = tmp + now
		}
		now = tmp
	}

	fmt.Println(min * m)
}
