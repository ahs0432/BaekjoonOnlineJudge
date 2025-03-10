package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var o1, o2 int
	fmt.Fscanln(reader, &o1, &o2)

	t := o1 % 30
	if t*12 == o2 {
		fmt.Println("O")
	} else {
		fmt.Println("X")
	}
}
