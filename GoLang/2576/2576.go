package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum, min := 0, 100
	var tmp int

	for i := 0; i < 7; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp%2 == 1 {
			sum += tmp
			if tmp < min {
				min = tmp
			}
		}
	}

	if min == 100 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
}
