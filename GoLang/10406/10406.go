package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w, n, p int
	fmt.Fscanln(reader, &w, &n, &p)

	count := 0
	var tmp int
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &tmp)

		if w <= tmp && tmp <= n {
			count++
		}
	}
	fmt.Println(count)
}
