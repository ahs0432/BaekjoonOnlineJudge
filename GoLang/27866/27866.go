package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	var n int
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &n)
	fmt.Println(string(s[n-1]))
}
