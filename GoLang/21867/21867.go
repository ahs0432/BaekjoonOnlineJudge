package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	s = strings.ReplaceAll(s, "J", "")
	s = strings.ReplaceAll(s, "A", "")
	s = strings.ReplaceAll(s, "V", "")

	if s == "" {
		fmt.Println("nojava")
	} else {
		fmt.Println(s)
	}
}
