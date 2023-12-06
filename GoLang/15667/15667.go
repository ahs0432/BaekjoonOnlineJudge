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

	for i := 1; i <= n; i++ {
		if i*i+i+1 == n {
			fmt.Println(i)
			break
		}
	}
}
