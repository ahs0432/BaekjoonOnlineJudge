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

	d := a + 7*b

	if d >= 1 && d <= 30 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
