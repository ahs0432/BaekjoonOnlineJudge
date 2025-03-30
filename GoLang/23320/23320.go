package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	var x, y int
	fmt.Fscanln(reader, &x, &y)

	aSum := n * x / 100
	bSum := 0

	for i := 0; i < n; i++ {
		if a[i] >= y {
			bSum++
		}
	}
	fmt.Println(aSum, bSum)
}
