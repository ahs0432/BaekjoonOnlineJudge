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

	table := make([]string, n)

	rgbTable := make([][]bool, n)
	rbTable := make([][]bool, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
		rgbTable[i] = make([]bool, n)
		rbTable[i] = make([]bool, n)
	}

	rgb := 0
	rb := 0

	qRGB := [][]int{}
	qRB := [][]int{}

	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !rgbTable[i][j] {
				rgb++
				qRGB = append(qRGB, []int{i, j, int(table[i][j])})
				rgbTable[i][j] = true
			}

			if !rbTable[i][j] {
				rb++
				qRB = append(qRB, []int{i, j, int(table[i][j])})
				rbTable[i][j] = true
			}

			for len(qRGB) != 0 {
				x := qRGB[0][0]
				y := qRGB[0][1]
				str := qRGB[0][2]
				qRGB = qRGB[1:]

				for _, move := range moves {
					mX := move[0] + x
					mY := move[1] + y

					if mX < 0 || mX == n || mY < 0 || mY == n || rgbTable[mX][mY] {
						continue
					}

					if str == int(table[mX][mY]) {
						qRGB = append(qRGB, []int{mX, mY, int(table[mX][mY])})
						rgbTable[mX][mY] = true
					}
				}
			}

			for len(qRB) != 0 {
				x := qRB[0][0]
				y := qRB[0][1]
				str := qRB[0][2]
				qRB = qRB[1:]

				for _, move := range moves {
					mX := move[0] + x
					mY := move[1] + y

					if mX < 0 || mX == n || mY < 0 || mY == n || rbTable[mX][mY] {
						continue
					}

					if (str == int(table[mX][mY])) || ((str == 71 || str == 82) && (int(table[mX][mY]) == 71 || int(table[mX][mY]) == 82)) {
						qRB = append(qRB, []int{mX, mY, int(table[mX][mY])})
						rbTable[mX][mY] = true
					}
				}
			}
		}
	}

	fmt.Println(rgb, rb)
}
