package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	if a/2 < b {
		fmt.Println(a / 2)
	} else {
		fmt.Println(b)
	}
}
