package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, tmp int
	fmt.Fscanln(reader, &n)

	count := make([]int, 53)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		count[tmp]++
	}

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	check := true
	for _, c := range s {
		tmp = 0
		if 'A' <= c && c <= 'Z' {
			tmp = int(c - 'A' + 1)
		} else if 'a' <= c && c <= 'z' {
			tmp = int(c - 'a' + 27)
		}

		count[tmp]--
		if count[tmp] < 0 {
			check = false
		}
	}

	if check {
		fmt.Println("y")
	} else {
		fmt.Println("n")
	}
}
