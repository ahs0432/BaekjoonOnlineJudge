package main

import (
	"bufio"
	"fmt"
	"os"
)

func clearTable(table [][]int, x1, x2, y1, y2 int) {
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			table[y][x] = 0
		}
	}
}

func cutTable(table [][]int, x1, x2, y1, y2 int) {
	if x1 == x2-1 && y1 == y2-1 {
		return
	}

	total := 0
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			total += table[y][x]
		}
	}

	nowLen := y2 - y1
	cutList := [][]int{
		{x1, x2 - (nowLen / 2), y1, y2 - (nowLen / 2)},
		{x1 + (nowLen / 2), x2, y1, y2 - (nowLen / 2)},
		{x1, x2 - (nowLen / 2), y1 + (nowLen / 2), y2},
		{x1 + (nowLen / 2), x2, y1 + (nowLen / 2), y2},
	}

	clearTarget := total % 4
	clearTable(table, cutList[clearTarget][0], cutList[clearTarget][1], cutList[clearTarget][2], cutList[clearTarget][3])
	cutList = append(cutList[:clearTarget], cutList[clearTarget+1:]...)

	for _, cut := range cutList {
		cutTable(table, cut[0], cut[1], cut[2], cut[3])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, ans int
	var tmp string
	var table [][]int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		table = make([][]int, n)
		for y := 0; y < n; y++ {
			fmt.Fscanln(reader, &tmp)
			table[y] = make([]int, n)

			for x := 0; x < n; x++ {
				table[y][x] = int(tmp[x] - '0')
			}
		}

		cutTable(table, 0, n, 0, n)

		ans = 0
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				ans += table[y][x]
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
