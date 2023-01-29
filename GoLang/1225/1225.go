package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a, &b)
	tmpB := b

	sum := int64(0)

	for a != "" {
		nowA := a[0]
		a = a[1:]

		for b != "" {
			nowB := b[0]
			b = b[1:]
			sum += int64((int(nowA) - 48) * (int(nowB) - 48))
		}
		b = tmpB
	}

	fmt.Println(sum)
}
