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

	lastWin := "D"
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &b[i])

		if a[i] > b[i] {
			lastWin = "A"
			aS += 3
		} else if a[i] < b[i] {
			lastWin = "B"
			bS += 3
		} else {
			aS++
			bS++
		}
	}

	fmt.Println(aS, bS)
	if aS > bS {
		fmt.Println("A")
	} else if aS < bS {
		fmt.Println("B")
	} else {
		fmt.Println(lastWin)
	}
}
