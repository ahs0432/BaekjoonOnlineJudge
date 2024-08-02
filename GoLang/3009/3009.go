package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	x := make([]int, 3)
	y := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &x[i], &y[i])
	}

	if x[0] == x[1] {
		fmt.Print(x[2], " ")
	} else if x[1] == x[2] {
		fmt.Print(x[0], " ")
	} else {
		fmt.Print(x[1], " ")
	}

	if y[0] == y[1] {
		fmt.Println(y[2])
	} else if y[1] == y[2] {
		fmt.Println(y[0])
	} else {
		fmt.Println(y[1])
	}
}
