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
	fmt.Println(a * b)
}
