package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	for i := 0; i < 5; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		if tmp < 40 {
			tmp = 40
		}
		sum += tmp
	}

	fmt.Println(sum / 5)
}
