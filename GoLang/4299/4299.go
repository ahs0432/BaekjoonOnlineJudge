package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	tmpB := (n - m) / 2
	tmpA := tmpB + m
	if tmpA >= 0 && tmpB >= 0 && (tmpA+tmpB) == n && (tmpA-tmpB) == m {
		fmt.Println(tmpA, tmpB)
	} else {
		fmt.Println(-1)
	}
}
