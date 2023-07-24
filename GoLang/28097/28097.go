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

	sum := -8
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(reader, &t)
		sum += (t + 8)
	}
	fmt.Println(sum/24, sum%24)
}
