package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	trees := make([]int, n)
	total := 0

	high := 0
	low := 0

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &trees[i])
		} else {
			fmt.Fscan(reader, &trees[i])
		}
		total += trees[i]

		if high < trees[i] {
			high = trees[i]
		}
	}

	answer := -1

	for low <= high {
		mid := (high + low) / 2
		count := int64(0)

		for _, t := range trees {
			if (t - mid) > 0 {
				count += int64(t - mid)
			}
		}

		if count == int64(m) {
			answer = mid
			break
		} else if count > int64(m) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if answer == -1 {
		answer = high
	}

	fmt.Println(answer)
}
