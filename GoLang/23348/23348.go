package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	level := make([]int, 3)
	var n, score int
	fmt.Fscanln(reader, &level[0], &level[1], &level[2])
	fmt.Fscanln(reader, &n)

	scores := make([]int, 0)
	tmp := make([]int, 3)
	for i := 0; i < n; i++ {
		score = 0
		for j := 0; j < 3; j++ {
			fmt.Fscanln(reader, &tmp[0], &tmp[1], &tmp[2])
			for k := 0; k < 3; k++ {
				score += tmp[k] * level[k]
			}
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	fmt.Println(scores[len(scores)-1])
}
