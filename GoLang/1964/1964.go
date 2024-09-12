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

	ans, p := 5, 7
	for i := 0; i < n-1; i++ {
		ans += p
		p += 3
		ans %= 45678
	}
	fmt.Println(ans)
}
