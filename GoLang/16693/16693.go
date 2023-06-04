package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a1, r1, p1, p2 int
	fmt.Fscanln(reader, &a1, &p1)
	fmt.Fscanln(reader, &r1, &p2)

	whole := float64(r1) * float64(r1) * 3.14159265359 / float64(p2)
	slice := float64(a1) / float64(p1)

	if whole > slice {
		fmt.Println("Whole pizza")
	} else {
		fmt.Println("Slice of pizza")
	}
}
