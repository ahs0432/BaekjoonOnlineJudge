package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscanln(reader, &n)

	fmt.Printf("%.6f\n", n*n*math.Pi)
	fmt.Printf("%.6f\n", 2*n*n)
}
