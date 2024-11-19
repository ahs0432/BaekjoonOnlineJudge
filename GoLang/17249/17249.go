package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	t := strings.Split(s, "(^0^)")

	l := 0
	r := 0
	for _, c := range t[0] {
		if c == '@' {
			l++
		}
	}

	for _, c := range t[1] {
		if c == '@' {
			r++
		}
	}

	fmt.Println(l, r)
}
