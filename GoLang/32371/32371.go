package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	key := make([][]rune, 4)
	visit := [4][10]bool{}
	dx := [9]int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := [9]int{-1, -1, -1, 0, 0, 0, 1, 1, 1}

	for i := 0; i < 4; i++ {
		line, _, _ := reader.ReadLine()
		key[i] = []rune(string(line))
	}

	strLine, _, _ := reader.ReadLine()
	str := []rune(string(strLine))

	for i := 0; i < len(str) && i < 9; i++ {
		found := false
		for j := 0; j < 4; j++ {
			for k := 0; k < 10; k++ {
				if key[j][k] == str[i] {
					visit[j][k] = true
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	var result rune
Loop:
	for i := 1; i < 3; i++ {
		for j := 1; j < 9; j++ {
			check := true
			for k := 0; k < 9; k++ {
				if !visit[i+dy[k]][j+dx[k]] {
					check = false
					break
				}
			}
			if check {
				result = key[i][j]
				break Loop
			}
		}
	}

	if result != 0 {
		fmt.Println(string(result))
	}
}
