package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	topingCounts1 := []float64{0.5, 0.5, 0.25, 0.0625, 0.5625}
	topingCounts2 := []float64{1, 30, 25, 10}

	topingInput1 := make([]float64, 5)
	topingInput2 := make([]float64, 4)

	var min, now, count int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader)

		count = 0
		min = 1000000000

		for i := 0; i < 5; i++ {
			fmt.Fscan(reader, &topingInput1[i])
			now = int(topingInput1[i] / topingCounts1[i])
			if min > now {
				min = now
			}
		}

		for i := 0; i < 4; i++ {
			fmt.Fscan(reader, &topingInput2[i])
			count += int(topingInput2[i] / topingCounts2[i])
		}

		if min < count {
			fmt.Fprintf(writer, "%d\n", min)
		} else {
			fmt.Fprintf(writer, "%d\n", count)
		}
	}
	writer.Flush()
}
