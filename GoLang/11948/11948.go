package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	min := 100
	for i := 0; i < 4; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)

		if min > tmp {
			min = tmp
		}
		sum += tmp
	}

	sum -= min

	top := 0
	for i := 0; i < 2; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)

		if top < tmp {
			top = tmp
		}
	}
	sum += top

	fmt.Println(sum)
}
