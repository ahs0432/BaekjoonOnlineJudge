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

	res := 1
	for i := 0; i < n; i++ {
		res *= 2
	}
	fmt.Println(res)
}
