package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := 0
	for i := 0; i < 10; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		t = (t + tmp) % 4
	}

	if t == 0 {
		fmt.Println("N")
	} else if t == 1 {
		fmt.Println("E")
	} else if t == 2 {
		fmt.Println("S")
	} else {
		fmt.Println("W")
	}
}
