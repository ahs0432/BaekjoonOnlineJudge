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
	hay := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &hay[i])
		sum += hay[i]
	}

	avg := sum / n

	res := 0

	for i := 0; i < n; i++ {
		if hay[i] > avg {
			res += (hay[i] - avg)
		}
	}
	fmt.Println(res)
}
