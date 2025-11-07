package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func zeroOut(value float64) float64 {
	if math.Abs(value) < 1e-9 {
		return 0.0
	}
	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, i, j, k float64
	fmt.Fscanln(reader, &a, &b, &c)
	fmt.Fscanln(reader, &i, &j, &k)
	tmp := math.Min(math.Min(a/i, b/j), c/k)
	fmt.Printf("%.10f %.10f %.10f\n", zeroOut(a-i*tmp), zeroOut(b-j*tmp), zeroOut(c-k*tmp))
}
