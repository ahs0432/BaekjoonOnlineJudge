package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n)

	table := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	var sum, avg, count int
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum, avg = 0, 0
			for k = i; k <= j; k++ {
				sum += table[k]
			}

			avg = sum / (j - i + 1)
			if sum%(j-i+1) != 0 {
				continue
			}

			for k = i; k <= j; k++ {
				if table[k] == avg {
					break
				}
			}

			if k != j+1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
