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

	if a-b > a+b {
		fmt.Println(a - b)
		fmt.Println(a + b)
	} else {
		fmt.Println(a + b)
		fmt.Println(a - b)
	}
}
