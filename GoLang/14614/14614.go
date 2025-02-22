package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	var c string
	fmt.Fscanln(reader, &a, &b, &c)

	if (c[len(c)-1]-0)%2 == 0 {
		fmt.Println(a)
	} else {
		fmt.Println(a ^ b)
	}
}
