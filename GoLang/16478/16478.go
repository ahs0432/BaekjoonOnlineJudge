package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var ab, bc, cd float64
	fmt.Fscanln(reader, &ab, &bc, &cd)
	fmt.Printf("%.6f\n", ab*cd/bc)
}
