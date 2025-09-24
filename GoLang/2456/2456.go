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

	s := make([]int, 3)
	score := make([]int, 3)
	count2 := make([]int, 3)
	count3 := make([]int, 3)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s[0], &s[1], &s[2])

		for j := 0; j < 3; j++ {
			score[j] += s[j]
			switch s[j] {
			case 3:
				count3[j]++
			case 2:
				count2[j]++
			}
		}
	}

	maxScore := 0
	maxCount2 := 0
	maxCount3 := 0
	id := -1

	for i := 0; i < 3; i++ {
		if maxScore < score[i] {
			maxScore = score[i]
			maxCount2 = count2[i]
			maxCount3 = count3[i]
			id = i
		} else if maxScore == score[i] {
			if maxCount3 < count3[i] {
				maxCount2 = count3[i]
				maxCount3 = count3[i]
				id = i
			} else if maxCount3 == count3[i] {
				if maxCount2 < count2[i] {
					maxCount2 = count2[i]
					id = i
				} else if maxCount2 == count2[i] {
					id = -1
				}
			}
		}
	}
	if id == -1 {
		fmt.Println(0, maxScore)
	} else {
		fmt.Println(id+1, maxScore)
	}
}
