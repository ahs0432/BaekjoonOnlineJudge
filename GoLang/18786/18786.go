package main

import (
	"bufio"
	"fmt"
	"os"
)

func area(x1, y1, x2, y2, x3, y3 int) int {
	xMatch := (x1-x2)*(x2-x3)*(x1-x3) == 0
	yMatch := (y1-y2)*(y1-y3)*(y2-y3) == 0

	if xMatch && yMatch {
		val := x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)
		if val < 0 {
			return -val
		}
		return val
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &arr[i][0], &arr[i][1])
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				currentArea := area(
					arr[i][0], arr[i][1],
					arr[j][0], arr[j][1],
					arr[k][0], arr[k][1],
				)
				if currentArea > maxArea {
					maxArea = currentArea
				}
			}
		}
	}

	fmt.Println(maxArea)
}
