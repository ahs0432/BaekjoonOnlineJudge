package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := make([]bool, 30)
	for i := 0; i < 28; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		n[tmp-1] = true
	}

	for i := 0; i < 30; i++ {
		if !n[i] {
			fmt.Println(i + 1)
		}

	}
}
