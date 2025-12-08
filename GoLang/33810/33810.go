package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	ori := "SciComLove"
	count := 0
	for i := 0; i < 10; i++ {
		if ori[i] != s[i] {
			count++
		}
	}
	fmt.Println(count)
}
