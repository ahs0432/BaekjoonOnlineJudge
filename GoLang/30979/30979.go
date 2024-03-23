package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t, n int
	fmt.Fscanln(reader, &t)
	fmt.Fscanln(reader, &n)

	total := 0
	var f int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &f)
		total += f
	}

	if t <= total {
		fmt.Println("Padaeng_i Happy")
	} else {
		fmt.Println("Padaeng_i Cry")
	}
}
