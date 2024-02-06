package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	tmp := (a + 1) * a / 2
	tmp %= 14579
	ans := tmp
	for i := a + 1; i <= b; i++ {
		tmp += i
		ans *= tmp
		ans %= 14579
	}
	fmt.Println(ans)
}
