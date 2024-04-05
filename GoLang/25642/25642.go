package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	for {
		b += a

		if b >= 5 {
			fmt.Println("yt")
			break
		}

		a += b
		if a >= 5 {
			fmt.Println("yj")
			break
		}
	}
}
