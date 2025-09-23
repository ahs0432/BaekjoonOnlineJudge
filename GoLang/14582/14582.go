package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	w := make([]int, 9)
	s := make([]int, 9)
	fmt.Fscanln(reader, &w[0], &w[1], &w[2], &w[3], &w[4], &w[5], &w[6], &w[7], &w[8])
	fmt.Fscanln(reader, &s[0], &s[1], &s[2], &s[3], &s[4], &s[5], &s[6], &s[7], &s[8])

	var win int
	var wScore int
	var sScore int

	for i := 0; i < 9; i++ {
		wScore += w[i]
		if wScore > sScore {
			win = 1
		}
		sScore += s[i]
	}

	if wScore < sScore && win == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
