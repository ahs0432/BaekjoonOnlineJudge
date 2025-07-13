package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	positions := make([][]float64, n)
	for i := 0; i < n; i++ {
		positions[i] = make([]float64, 2)
		fmt.Fscanln(reader, &positions[i][0], &positions[i][1])
	}

	var m, p int
	fmt.Fscanln(reader, &m)

	var tmp []int
	var ans float64
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &p)

		ans = 0
		tmp = make([]int, p)
		for j := 0; j < p; j++ {
			if j != p-1 {
				fmt.Fscan(reader, &tmp[j])
			} else {
				fmt.Fscanln(reader, &tmp[j])
			}
		}

		for j := 1; j < p; j++ {
			ans += math.Sqrt(math.Pow(positions[tmp[j-1]][0]-positions[tmp[j]][0], 2) + math.Pow(positions[tmp[j-1]][1]-positions[tmp[j]][1], 2))
		}
		fmt.Println(math.Round(ans))
	}
	writer.Flush()
}
