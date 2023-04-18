package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < 3; i++ {
		start := make([]int, 3)
		end := make([]int, 3)

		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &start[j])
		}

		for j := 0; j < 3; j++ {
			if j == 2 {
				fmt.Fscanln(reader, &end[j])
			} else {
				fmt.Fscan(reader, &end[j])
			}
		}

		result := make([]int, 3)
		if end[2]-start[2] < 0 {
			result[2] = 60 + (end[2] - start[2])
			start[1] += 1
		} else {
			result[2] = (end[2] - start[2])
		}

		if end[1]-start[1] < 0 {
			result[1] = 60 + (end[1] - start[1])
			start[0] += 1
		} else {
			result[1] = (end[1] - start[1])
		}

		result[0] = end[0] - start[0]

		fmt.Fprintln(writer, result[0], result[1], result[2])
	}
	writer.Flush()
}
