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

	ans := 100.0 / ((1.609344 / 3.785411784) * n)
	fmt.Printf("%.6f\n", ans)
}
