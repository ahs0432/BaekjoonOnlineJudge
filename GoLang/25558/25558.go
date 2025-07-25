package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(i int64) int64 {
	if i >= 0 {
		return i
	}
	return -i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)

	var sx, sy, ex, ey int64
	fmt.Fscanln(reader, &sx, &sy, &ex, &ey)

	x := make([]int64, 100)
	y := make([]int64, 100)

	var tmp int64
	tmpMax := int64(1000000000000000007)
	ans := 0
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &m)

		tmp = 0
		for j := 0; j < m; j++ {
			fmt.Fscanln(reader, &x[j], &y[j])
		}

		for j := 1; j < m; j++ {
			tmp += abs(x[j]-x[j-1]) + abs(y[j]-y[j-1])
		}

		tmp += abs(x[0]-sx) + abs(y[0]-sy) + abs(x[m-1]-ex) + abs(y[m-1]-ey)

		if tmp < tmpMax {
			ans = i
			tmpMax = tmp
		}
	}

	fmt.Println(ans)
}
