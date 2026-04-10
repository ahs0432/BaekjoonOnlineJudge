package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, nq, p int
	fmt.Fscanf(reader, "%d %d %d\n", &n, &nq, &p)

	answers := make([][]int, n)
	for i := 0; i < n; i++ {
		answers[i] = make([]int, nq+1)
		for j := 1; j <= nq; j++ {
			fmt.Fscanf(reader, "%d", &answers[i][j])
		}
		fmt.Fscanln(reader)
	}

	type condition struct {
		qIdx int
		val  int
	}
	ps := make([]condition, p)
	for i := 0; i < p; i++ {
		fmt.Fscanf(reader, "%d %d\n", &ps[i].qIdx, &ps[i].val)
	}

	count := 0

	for _, ans := range answers {
		isOk := true
		for _, cond := range ps {
			if ans[cond.qIdx] != cond.val {
				isOk = false
				break
			}
		}

		if isOk {
			count++
		}
	}

	fmt.Println(count)
}
