package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var x, y, blocks int

	fmt.Fscanln(reader, &x, &y, &blocks)

	table := make([][]int, x)

	min := 257
	max := 0

	for i := 0; i < x; i++ {
		table[i] = make([]int, y)
		for j := 0; j < y; j++ {
			if j == y-1 {
				fmt.Fscanln(reader, &table[i][j])
			} else {
				fmt.Fscan(reader, &table[i][j])
			}

			if max < table[i][j] {
				max = table[i][j]
			}

			if min > table[i][j] {
				min = table[i][j]
			}
		}
	}

	time := -1
	height := 0

	for ; min <= max; min++ {
		remove := 0
		stacking := 0

		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				if min == table[i][j] {
					continue
				} else if min < table[i][j] {
					remove += (table[i][j] - min)
				} else {
					stacking += (min - table[i][j])
				}
			}
		}

		if stacking <= (blocks + remove) {
			tmpTime := (remove * 2) + (stacking)
			if time == -1 || tmpTime < time {
				time = tmpTime
				height = min
			} else if tmpTime == time && height < min {
				time = tmpTime
				height = min
			}
		}
	}

	fmt.Println(time, height)
}
