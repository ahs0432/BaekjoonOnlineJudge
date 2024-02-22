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

	var score int
	t := make([]float64, 7)
	for i := 0; i < n; i++ {
		score = 0
		for j := 0; j < 7; j++ {
			fmt.Fscan(reader, &t[j])
		}

		score += int(9.23076 * math.Pow(26.7-t[0], 1.835))
		score += int(4.99087 * math.Pow(42.5-t[3], 1.81))
		score += int(0.11193 * math.Pow(254-t[6], 1.88))
		score += int(1.84523 * math.Pow(t[1]-75, 1.348))
		score += int(56.0211 * math.Pow(t[2]-1.5, 1.05))
		score += int(0.188807 * math.Pow(t[4]-210, 1.41))
		score += int(15.9803 * math.Pow(t[5]-3.8, 1.04))

		fmt.Fprintln(writer, score)
	}
	writer.Flush()
}
