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

	ans := ""
	for i := 1; i <= len(s); i++ {
		ans = string(s[len(s)-i]) + ans
		if i%3 == 0 && len(s) != i {
			ans = "," + ans
		}
	}
	fmt.Println(ans)
}
