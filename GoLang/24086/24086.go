package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	if b < a {
		fmt.Println(a - b)
	} else {
		fmt.Println(b - a)
	}
}
