package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscanln(reader, &m, &n)
	if m > n {
		fmt.Println(n*2 - 1)
	} else {
		fmt.Println(m*2 - 2)
	}
}
