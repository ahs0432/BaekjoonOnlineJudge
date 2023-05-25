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

	sum := 0

	for i := 0; i < 3; i++ {
		var tmp int

		if i == 2 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}

		if tmp < n {
			sum += tmp
		} else {
			sum += n
		}
	}
	fmt.Println(sum)
}
