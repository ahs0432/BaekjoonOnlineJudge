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
	fmt.Printf("%.8f\n", math.Sqrt(n)*4)
}
