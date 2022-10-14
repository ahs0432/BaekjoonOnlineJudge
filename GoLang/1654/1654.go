package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, n int
	fmt.Fscanln(reader, &k, &n)

	table := make([]int, k)
	high := 0
	low := 1

	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &table[i])

		if high < table[i] {
			high = table[i]
		}
	}

	answer := 0

	for low <= high {
		mid := (high + low) / 2
		count := 0

		for i := 0; i < k; i++ {
			count += (table[i] / mid)
		}

		if count >= n {
			low = mid + 1
			if answer < mid {
				answer = mid
			}
		} else {
			high = mid - 1
		}
	}

	fmt.Println(answer)
}
