package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x1, y1, z1, x2, y2, z2 float64
	fmt.Fscanln(reader, &x1, &y1, &z1)
	fmt.Fscanln(reader, &x2, &y2, &z2)

	if x1 <= x2 && y1 <= y2 && z1 <= z2 {
		fmt.Println("A")
	} else if x1/2 <= x2 && y1 <= y2 && z1 <= z2 {
		fmt.Println("B")
	} else if y1 <= y2 && z1 <= z2 {
		fmt.Println("C")
	} else if y1/2 <= y2 && z1/2 <= z2 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
