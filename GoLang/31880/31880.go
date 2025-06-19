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

	sum := 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		sum += tmp
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp != 0 {
			sum *= tmp
		}
	}
	fmt.Println(sum)
}
