package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, w float64
	fmt.Fscanln(reader, &h, &w)

	minHW, maxHW := min(h, w), max(h, w)

	if maxHW/3 <= minHW {
		h = maxHW / 3
	} else {
		h = minHW
	}
	w = minHW / 2
	fmt.Printf("%f\n", max(h, w))
}
