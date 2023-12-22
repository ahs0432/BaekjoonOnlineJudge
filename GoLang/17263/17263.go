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

	max := 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if max < tmp {
			max = tmp
		}
	}
	fmt.Println(max)
}
