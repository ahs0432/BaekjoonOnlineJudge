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
	fmt.Println(a + b + c)
}
