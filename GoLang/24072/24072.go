package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	if a <= c && c < b {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
