package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	max := 0
	var tmp, run int
	trick := make([]int, 5)
	for i := 0; i < n; i++ {
		run = 0
		for j := 0; j < 2; j++ {
			fmt.Fscan(reader, &tmp)

			if tmp > run {
				run = tmp
			}
		}

		for j := 0; j < 5; j++ {
			fmt.Fscan(reader, &trick[j])
		}

		sort.Slice(trick, func(k, l int) bool {
			return trick[k] > trick[l]
		})

		if max < run+trick[0]+trick[1] {
			max = run + trick[0] + trick[1]
		}
	}

	fmt.Println(max)
}
