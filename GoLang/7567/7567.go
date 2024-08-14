package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	ans := 10

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			ans += 5
		} else {
			ans += 10
		}
	}
	fmt.Println(ans)
}
