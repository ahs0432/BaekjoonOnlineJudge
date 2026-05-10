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

	var x, y int
	q1, q2, q3, q4, axis := 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)

		if x > 0 && y > 0 {
			q1 += 1
		} else if x < 0 && y > 0 {
			q2 += 1
		} else if x < 0 && y < 0 {
			q3 += 1
		} else if x > 0 && y < 0 {
			q4 += 1
		} else {
			axis++
		}
	}

	fmt.Println("Q1:", q1)
	fmt.Println("Q2:", q2)
	fmt.Println("Q3:", q3)
	fmt.Println("Q4:", q4)
	fmt.Println("AXIS:", axis)
}
