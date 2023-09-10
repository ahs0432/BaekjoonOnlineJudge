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
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp == 0 {
			sum -= 1
		} else {
			sum += 1
		}
	}

	if sum > 0 {
		fmt.Println("Junhee is cute!")
	} else {
		fmt.Println("Junhee is not cute!")
	}
}
