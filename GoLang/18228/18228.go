package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	minD := []int{1000000001, 1000000001}
	var tmp, id int

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp == -1 {
			id++
		} else {
			minD[id] = min(minD[id], tmp)
		}
	}
	fmt.Println(minD[0] + minD[1])
}
