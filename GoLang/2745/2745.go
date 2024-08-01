package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	var b int
	fmt.Fscanln(reader, &n, &b)

	ans := 0
	now := 1
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] >= 'A' && n[i] <= 'Z' {
			ans += (int(n[i]) - 55) * now
		} else {
			ans += (int(n[i]) - 48) * now
		}
		now *= b
	}
	fmt.Println(ans)
}
