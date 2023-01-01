package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	q := [][]int{{a, b, 1}}

	for len(q) != 0 {
		tmpA := q[0][0]
		tmpB := q[0][1]
		count := q[0][2]
		q = q[1:]

		if tmpA == tmpB {
			fmt.Println(count)
			return
		} else {
			if tmpA*10+1 <= tmpB {
				q = append(q, []int{tmpA*10 + 1, tmpB, count + 1})
			}

			if tmpA*2 <= tmpB {
				q = append(q, []int{tmpA * 2, tmpB, count + 1})
			}
		}
	}

	fmt.Println(-1)
}
