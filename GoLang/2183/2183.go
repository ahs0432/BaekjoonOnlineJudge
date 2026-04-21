package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n, &s)
	fmt.Printf("%c\n", s[len(s)-1])
}
