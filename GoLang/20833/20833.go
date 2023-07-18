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

	ans := 0
	for i := 1; i <= n; i++ {
		ans += (i * i * i)
	}
	fmt.Println(ans)
}
