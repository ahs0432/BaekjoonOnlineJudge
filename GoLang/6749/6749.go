package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2 int
	fmt.Fscanln(reader, &n1)
	fmt.Fscanln(reader, &n2)

	fmt.Println(n2 + (n2 - n1))
}
