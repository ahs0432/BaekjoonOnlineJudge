package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	tmp := ""
	for i := 0; i < len(s); i++ {
		tmp += string(s[len(s)-1-i])
	}
	return tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	var l int
	var s string
	var check bool
	for i := 0; i < n; i++ {
		s = reverse(table[i])

		for j := 0; j < n; j++ {
			if s == table[j] {
				check = true
				l = len(s)
				break
			}
		}

		if check {
			break
		}
	}
	fmt.Println(l, string(s[l/2]))
}
