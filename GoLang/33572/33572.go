package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int64
	fmt.Fscanln(reader, &n, &m)

	var sum, a int64
	for i := int64(0); i < n; i++ {
		fmt.Fscan(reader, &a)
		sum += a
	}

	if sum-n < m {
		fmt.Println("OUT")
	} else {
		fmt.Println("DIMI")
	}
}
