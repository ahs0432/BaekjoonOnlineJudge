package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	count, total := 0, 0
	for total < c {
		total += a
		count++
		if count%7 == 0 {
			total += b
		}
	}
	fmt.Println(count)
}
