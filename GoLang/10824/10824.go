package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d int64
	fmt.Fscanln(reader, &a, &b, &c, &d)

	now := int64(10)
	for b/now > 0 {
		now *= 10
	}

	sum1 := (a * now) + b

	now = int64(10)
	for d/now > 0 {
		now *= 10
	}

	sum2 := (c * now) + d

	fmt.Println(sum1 + sum2)
}
