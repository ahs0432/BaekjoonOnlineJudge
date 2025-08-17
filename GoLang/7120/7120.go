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

	ans := string(s[0])
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			ans = ans + string(s[i])
		}
	}
	fmt.Println(ans)
}
