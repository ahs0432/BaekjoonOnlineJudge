package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverse(s string) string {
	var tmp []rune
	for _, c := range s {
		tmp = append([]rune{c}, tmp...)
	}
	return string(tmp)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a, &b)

	a, b = reverse(a), reverse(b)
	var ans string

	minLen := len(a)
	checkLen := false
	diffLen := len(b) - len(a)

	if minLen > len(b) {
		minLen = len(b)
		checkLen = true
		diffLen = len(a) - len(b)
	}

	for i := 0; i < minLen; i++ {
		ans = strconv.Itoa(int(a[i]-48+b[i]-48)) + ans
	}

	for i := minLen; i < minLen+diffLen; i++ {
		if checkLen {
			ans = string(a[i]) + ans
		} else {
			ans = string(b[i]) + ans
		}
	}

	fmt.Println(ans)
}
