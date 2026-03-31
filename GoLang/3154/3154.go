package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanf(reader, "%02d:%02d", &a, &b)

	v := [10][10]int{
		{0, 4, 3, 4, 3, 2, 3, 2, 1, 2},
		{4, 0, 1, 2, 1, 2, 3, 2, 3, 4},
		{3, 1, 0, 1, 2, 1, 2, 3, 2, 3},
		{4, 2, 1, 0, 3, 2, 1, 4, 3, 2},
		{3, 1, 2, 3, 0, 1, 2, 1, 2, 3},
		{2, 2, 1, 2, 1, 0, 1, 2, 1, 2},
		{3, 3, 2, 1, 2, 1, 0, 3, 2, 1},
		{2, 2, 3, 4, 1, 2, 3, 0, 1, 2},
		{1, 3, 2, 3, 2, 1, 2, 1, 0, 1},
		{2, 4, 3, 2, 3, 2, 1, 2, 1, 0},
	}

	minSum := math.MaxInt64
	var resX, resY, sum int

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i%24 != a || j%60 != b {
				continue
			}

			sum = v[i/10][i%10] + v[i%10][j/10] + v[j/10][j%10]

			if sum < minSum {
				minSum = sum
				resX, resY = i, j
			}
		}
	}

	fmt.Printf("%02d:%02d\n", resX, resY)
}
