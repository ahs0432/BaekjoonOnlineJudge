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

	for strings.Contains(s, "PS4") || strings.Contains(s, "PS5") {
		s = strings.ReplaceAll(strings.ReplaceAll(s, "PS4", "PS"), "PS5", "PS")
	}

	fmt.Println(s)
}
