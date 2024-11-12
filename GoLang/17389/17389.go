package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	ans, bonus := 0, 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] == 'O' {
			ans += i + bonus
			bonus++
		} else {
			bonus = 0
		}
	}
	fmt.Println(ans)
}
