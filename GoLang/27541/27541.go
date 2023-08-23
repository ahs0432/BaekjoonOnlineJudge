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
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	if s[n-1] == byte('G') {
		s = s[0 : n-1]
	} else {
		s = s + "G"
	}

	fmt.Println(s)
}
