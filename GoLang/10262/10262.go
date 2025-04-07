package main

import (
	"bufio"
	"fmt"
	"os"
)

func calcAVG(a, b float64) float64 {
	var ans float64 = 0.0
	for i := a; i <= b; i++ {
		ans += i
	}
	return ans / (b - a + 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	gunnar := [][]float64{{0, 0}, {0, 0}}
	emma := [][]float64{{0, 0}, {0, 0}}

	fmt.Fscanln(reader, &gunnar[0][0], &gunnar[0][1], &gunnar[1][0], &gunnar[1][1])
	fmt.Fscanln(reader, &emma[0][0], &emma[0][1], &emma[1][0], &emma[1][1])

	g := calcAVG(gunnar[0][0], gunnar[0][1]) + calcAVG(gunnar[1][0], gunnar[1][1])
	e := calcAVG(emma[0][0], emma[0][1]) + calcAVG(emma[1][0], emma[1][1])

	if g > e {
		fmt.Println("Gunnar")
	} else if g == e {
		fmt.Println("Tie")
	} else {
		fmt.Println("Emma")
	}
}
