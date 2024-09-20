package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := make([]int, 10)
	b := make([]int, 10)

	aS := 0
	bS := 0

	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &b[i])
		if a[i] > b[i] {
			aS++
		} else if a[i] < b[i] {
			bS++
		}
	}

	if aS > bS {
		fmt.Println("A")
	} else if aS < bS {
		fmt.Println("B")
	} else {
		fmt.Println("D")
	}
}
