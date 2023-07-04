package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, z int
	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &y)
	fmt.Fscanln(reader, &z)

	if x+y <= z {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
