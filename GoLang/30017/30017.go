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

	for a >= 2 {
		if b >= a-1 {
			fmt.Println(2*a - 1)
			break
		} else {
			a--
		}
	}
}
