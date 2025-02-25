package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	scoreTable := []int{10, 8, 6, 5, 4, 3, 2, 1, 0}
	timeTable := make([][]string, 8)
	rScore, bScore := 0, 0

	for i := 0; i < 8; i++ {
		timeTable[i] = make([]string, 2)
		fmt.Fscanln(reader, &timeTable[i][0], &timeTable[i][1])
	}

	sort.Slice(timeTable, func(i, j int) bool {
		return timeTable[i][0] < timeTable[j][0]
	})

	for i := 0; i < 8; i++ {
		if timeTable[i][1] == "R" {
			rScore += scoreTable[i]
		} else {
			bScore += scoreTable[i]
		}
	}

	if rScore > bScore {
		fmt.Println("Red")
	} else if rScore == bScore {
		if timeTable[0][1] == "R" {
			fmt.Println("Red")
		} else {
			fmt.Println("Blue")
		}
	} else {
		fmt.Println("Blue")
	}
}
