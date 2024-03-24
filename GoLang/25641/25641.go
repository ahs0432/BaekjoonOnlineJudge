package main

import (
	"bufio"
	"fmt"
	"os"
)

func diff(s string) bool {
	tt := 0
	so := 0
	for _, c := range s {
		if c == 't' {
			tt++
		} else {
			so++
		}
	}

	return tt == so
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	for {
		if len(s)%2 != 0 || !diff(s) {
			s = s[1:]
		} else {
			break
		}
	}
	fmt.Println(s)
}
