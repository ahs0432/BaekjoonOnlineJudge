package main

import (
	"bufio"
	"fmt"
	"os"
)

var table [][]int

func checkS(start, end []int) (int, int) {
	if start[0] == end[0] && start[1] == end[1] {
		if table[start[0]][start[1]] == 1 {
			return 0, 1
		} else {
			return 1, 0
		}
	}

	num := -1
	check := true

	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			if num == -1 {
				num = table[i][j]
			} else if num != table[i][j] {
				check = false
				break
			}
		}
	}

	if check {
		if num == 1 {
			return 0, 1
		} else {
			return 1, 0
		}
	} else {
		white, blue := 0, 0
		plusData := ((end[0] - start[0] + 1) / 2)
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				tmpWhite, tmpBlue := checkS([]int{start[0] + (plusData * i), start[1] + (plusData * j)}, []int{start[0] + (plusData * (i + 1)) - 1, start[1] + (plusData * (j + 1)) - 1})
				white += tmpWhite
				blue += tmpBlue
			}
		}
		return white, blue
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table = make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j == n {
				fmt.Fscanln(reader, &table[i][j])
			} else {
				fmt.Fscan(reader, &table[i][j])
			}
		}
	}

	white, blue := checkS([]int{0, 0}, []int{n - 1, n - 1})

	fmt.Println(white)
	fmt.Println(blue)
}
