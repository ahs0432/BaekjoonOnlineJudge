package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, r int
	fmt.Fscanln(reader, &l)
	fmt.Fscanln(reader, &r)

	n, sum, t := l, 0, 2
	for {
		n = int(float64(n) * (float64(r) / 100.0))
		if n <= 5 {
			break
		}
		sum += t * n
		t *= 2
	}
	fmt.Println(sum)
}
