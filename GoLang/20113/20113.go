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

	var tmp int
	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp == 0 {
			continue
		}
		count[tmp]++
	}

	max := 0
	maxCount := 0
	for i := 1; i < n+1; i++ {
		if count[max] < count[i] {
			max = i
			maxCount = 1
		} else if count[max] == count[i] {
			maxCount++
		}
	}

	if maxCount >= 2 || max == 0 {
		fmt.Println("skipped")
	} else {
		fmt.Println(max)
	}
}
