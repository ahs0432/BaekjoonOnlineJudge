package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscanln(reader, &n)
	fmt.Printf("%.0f\n", (n/2)*(n/2))
}
