package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, max, sum := 0, 0, 0
	var tmp int

	for i := 0; i < 5; i++ {
		sum = 0

		for j := 0; j < 4; j++ {
			fmt.Fscan(reader, &tmp)
			sum += tmp
		}

		if sum > max {
			a = i
			max = sum
		}
	}

	fmt.Println(a+1, max)
}
