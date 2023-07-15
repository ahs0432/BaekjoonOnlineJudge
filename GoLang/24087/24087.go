package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, a, b int
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	ans := 250

	if s-a > 0 {
		ans += ((s - a) / b) * 100
		if (s-a)%b != 0 {
			ans += 100
		}
	}

	fmt.Println(ans)
}
