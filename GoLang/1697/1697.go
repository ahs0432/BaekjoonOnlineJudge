package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	already := map[int]bool{}
	queue := [][]int{{n, 0}}
	already[n] = true

	for len(queue) != 0 {
		if queue[0][0] == m {
			fmt.Println(queue[0][1])
			break
		}

		if queue[0][0] < m {
			if !already[queue[0][0]*2] {
				queue = append(queue, []int{queue[0][0] * 2, queue[0][1] + 1})
				already[queue[0][0]*2] = true
			}

			if !already[queue[0][0]+1] {
				queue = append(queue, []int{queue[0][0] + 1, queue[0][1] + 1})
				already[queue[0][0]+1] = true
			}
		}

		if queue[0][0] > 0 && !already[queue[0][0]-1] {
			queue = append(queue, []int{queue[0][0] - 1, queue[0][1] + 1})
			already[queue[0][0]-1] = true
		}

		queue = queue[1:]
	}

}
