package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var c1, b1, p1 int
	var c2, b2, p2 int

	fmt.Fscanln(reader, &c1, &b1, &p1)
	fmt.Fscanln(reader, &c2, &b2, &p2)

	sum := 0
	if c2-c1 > 0 {
		sum += c2 - c1
	}

	if b2-b1 > 0 {
		sum += b2 - b1
	}

	if p2-p1 > 0 {
		sum += p2 - p1
	}

	fmt.Println(sum)
}
