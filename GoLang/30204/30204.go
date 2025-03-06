package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscanln(reader, &n, &x)

	var tmp int
	var sum int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		sum += tmp
	}

	if sum%x == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
