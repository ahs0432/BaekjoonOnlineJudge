package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, h, v int
	fmt.Fscanln(reader, &n, &h, &v)
	res := max(h, n-h) * max(v, n-v)
	fmt.Println(res * 4)
}
