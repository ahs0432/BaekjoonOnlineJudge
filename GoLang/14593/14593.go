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

	players := make([][]int, n)

	for i := 0; i < n; i++ {
		players[i] = make([]int, 4)
		players[i][3] = i + 1
		fmt.Fscanln(reader, &players[i][0], &players[i][1], &players[i][2])
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i][0] != players[j][0] {
			return players[i][0] > players[j][0]
		} else {
			if players[i][1] != players[j][1] {
				return players[i][1] < players[j][1]
			} else {
				return players[i][2] < players[j][2]
			}
		}
	})
	fmt.Println(players[0][3])
}
