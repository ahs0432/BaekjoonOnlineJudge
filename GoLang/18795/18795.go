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

	sum := int64(0)
	var tmp int64
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		sum += tmp
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &tmp)
		sum += tmp
	}
	fmt.Println(sum)
}
