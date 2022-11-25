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

	max := 0
	table := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &table[i])
		} else {
			fmt.Fscan(reader, &table[i])
		}

		if max < table[i] {
			max = table[i]
		}
	}

	prime := make([]bool, max+1)
	prime[0] = true
	prime[1] = true

	for i := 2; i <= max/2; i++ {
		for j := i * 2; j <= max; j += i {
			prime[j] = true
		}
	}

	answer := 0
	for _, t := range table {
		if !prime[t] {
			answer++
		}
	}

	fmt.Println(answer)
}
