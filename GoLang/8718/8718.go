package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, k int
	fmt.Fscanln(reader, &x, &k)

	max := 0

	if 7000*k <= 1000*x {
		max = 7000 * k
	}

	if 3500*k <= 1000*x && max < 3500*k {
		max = 3500 * k
	}
	if 1750*k <= 1000*x && max < 1750*k {
		max = 1750 * k
	}

	fmt.Println(max)
}
