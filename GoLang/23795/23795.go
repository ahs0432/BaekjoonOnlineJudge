package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	for {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		if tmp == -1 {
			break
		}

		sum += tmp
	}

	fmt.Println(sum)
}
