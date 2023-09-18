package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, tmp int
	fmt.Fscanln(reader, &n)

	sum, now := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp == 1 {
			now++
			sum += now
		} else {
			now = 0
		}
	}

	fmt.Println(sum)
}
