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

	sum := 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		sum += (tmp - 1)
	}

	fmt.Println(sum + 1)
}
