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
	max, min := 0, 1000
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if max < tmp {
			max = tmp
		}

		if min > tmp {
			min = tmp
		}
	}
	fmt.Println(max - min)
}
