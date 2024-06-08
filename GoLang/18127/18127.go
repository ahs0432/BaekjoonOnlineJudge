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
	fmt.Println(((b + 1) * (2 + b*(a-2))) / 2)
}
