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

	if a == b && b == c {
		fmt.Println("*")
	} else if a == b && b != c {
		fmt.Println("C")
	} else if a == c && b != c {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
