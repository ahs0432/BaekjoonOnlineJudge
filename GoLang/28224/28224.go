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
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		sum += tmp
	}

	fmt.Println(sum)
}
