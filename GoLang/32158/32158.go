package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	var s string
	fmt.Fscan(reader, &s)

	runes := []rune(s)

	var pi []int
	var ci []int

	for i, c := range runes {
		if c == 'P' {
			pi = append(pi, i)
		} else if c == 'C' {
			ci = append(ci, i)
		}
	}

	minLen := len(pi)
	if len(ci) < minLen {
		minLen = len(ci)
	}

	for i := 0; i < minLen; i++ {
		pIdx := pi[i]
		cIdx := ci[i]
		runes[pIdx], runes[cIdx] = runes[cIdx], runes[pIdx]
	}

	fmt.Println(string(runes))
}
