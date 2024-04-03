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
		fmt.Fscan(reader, &tmp)
		sum += tmp
	}

	if sum%3 == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
