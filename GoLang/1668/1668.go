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
		fmt.Fscanln(reader, &table[i])
		if max < table[i] {
			max = table[i]
		}
	}

	lc, lm := 0, 0
	rc, rm := 0, 0
	for i := 0; i < n; i++ {
		if table[i] > lm {
			lm = table[i]
			lc++
		}

		if table[i] == max {
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if table[i] > rm {
			rm = table[i]
			rc++
		}

		if table[i] == max {
			break
		}
	}

	fmt.Println(lc)
	fmt.Println(rc)
}
