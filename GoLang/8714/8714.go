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

	var tmp int
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp != 0 {
			count++
		}
	}

	if count < n-count {
		fmt.Println(count)
	} else {
		fmt.Println(n - count)
	}
}
