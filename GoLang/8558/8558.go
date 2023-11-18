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

	ans := 1
	for i := 2; i <= n; i++ {
		ans = (ans * i) % 10
	}
	fmt.Println(ans)
}
