package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var m, res int
	for i := 0; i < n; i++ {
		res = 0
		fmt.Fscanln(reader, &m)
		missonList := make([][]int, m)
		for j := 0; j < m; j++ {
			missonList[j] = make([]int, 3)
			fmt.Fscanln(reader, &missonList[j][0], &missonList[j][1], &missonList[j][2])
		}
		kda := make([]int, 3)
		fmt.Fscanln(reader, &kda[0], &kda[1], &kda[2])

		for _, mission := range missonList {
			if score := (mission[0]*kda[0] - mission[1]*kda[1] + mission[2]*kda[2]); score > 0 {
				res += score
			}
		}

		fmt.Fprintln(writer, res)
	}
	writer.Flush()
}
