package main

import (
	"bufio"
	"fmt"
	"os"
)

var table [][]int

func quadSlice(min []int, max []int) string {
	now := -1
	answer := ""
	for i := min[0]; i < max[0]; i++ {
		check := true
		for j := min[1]; j < max[1]; j++ {
			if now == -1 {
				now = table[i][j]
			} else if now != table[i][j] {
				xPlus := (max[0] - min[0]) / 2
				yPlus := (max[1] - min[1]) / 2

				for x := 1; x <= 2; x++ {
					for y := 1; y <= 2; y++ {
						answer += quadSlice([]int{min[0] + (xPlus * (x - 1)), min[1] + (yPlus * (y - 1))}, []int{min[0] + (xPlus * x), min[1] + (yPlus * y)})
					}
				}
				check = false
				break
			}
		}

		if !check {
			break
		}
	}

	if answer == "" {
		return string(rune(now + 48))
	}

	return "(" + answer + ")"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table = make([][]int, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			table[i][j] = int(tmp[j] - 48)
		}
	}

	answer := quadSlice([]int{0, 0}, []int{n, n})
	fmt.Println(answer)
}
