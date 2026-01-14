package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a, b int
	fmt.Fscanln(reader, &n, &a, &b)

	ansA, ansB := 1, 1
	for i := 0; i < n; i++ {
		ansA += a
		ansB += b

		if ansA < ansB {
			ansA, ansB = ansB, ansA
		} else if ansA == ansB {
			ansB--
		}
	}
	fmt.Println(ansA, ansB)
}
