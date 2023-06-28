package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w, h int
	fmt.Fscanln(reader, &w, &h)
	fmt.Printf("%.9f\n", float64(w+h)-math.Sqrt(float64(w*w+h*h)))
}
