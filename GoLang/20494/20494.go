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
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		ans += len(s)
	}
	fmt.Println(ans)
}
