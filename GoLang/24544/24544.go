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

	ir := make([]int, n)

	sumA := 0
	sumB := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ir[i])
		sumA += ir[i]
	}

	var ap int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ap)
		if ap == 0 {
			sumB += ir[i]
		}
	}
	fmt.Println(sumA)
	fmt.Println(sumB)
}
