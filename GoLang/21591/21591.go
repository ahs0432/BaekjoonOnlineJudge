package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w1, h1, w2, h2 int
	fmt.Fscanln(reader, &w1, &h1, &w2, &h2)

	if w1 >= w2+2 && h1 >= h2+2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
