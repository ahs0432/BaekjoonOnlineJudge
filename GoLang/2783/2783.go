package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y float64
	fmt.Fscanln(reader, &x, &y)

	var n int
	fmt.Fscanln(reader, &n)

	min := x / y
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)
		if min > x/y {
			min = x / y
		}
	}
	fmt.Printf("%.2f\n", min*1000)
}
