package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p1, p2, q1, q2 int64
	fmt.Fscanln(reader, &p1, &q1, &p2, &q2)
	if p1*p2%(q1*q2*2) == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
