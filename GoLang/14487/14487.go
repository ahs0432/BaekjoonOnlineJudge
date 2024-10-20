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

	var sum, max, tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp > max {
			max = tmp
		}
		sum += tmp
	}
	fmt.Println(sum - max)
}
