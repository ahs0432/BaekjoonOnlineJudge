package main

import (
	"bufio"
	"fmt"
	"os"
)

func big(c byte) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	ans := int64(0)
	count := int64(0)
	for i := 0; i < len(s); i++ {
		if big(s[i]) {
			ans += (4 - count%4) % 4
			count = 0
		}
		count++
	}
	fmt.Println(ans)
}
