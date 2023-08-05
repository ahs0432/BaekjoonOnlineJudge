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

	if c < b {
		b, c = c, b
	}

	if c < a {
		a, c = c, a
	}

	if a+b == c {
		fmt.Println("1")
	} else if a*b == c {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}
