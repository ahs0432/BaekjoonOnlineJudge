package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	if n == m {
		if n == 0 {
			fmt.Println("Not a moose")
		} else {
			fmt.Println("Even", n+m)
		}
	} else if n > m {
		fmt.Println("Odd", n*2)
	} else {
		fmt.Println("Odd", m*2)
	}
}
