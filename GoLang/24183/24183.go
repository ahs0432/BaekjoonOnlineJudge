package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var c4, a3, a4 float64
	fmt.Fscanln(reader, &c4, &a3, &a4)

	envelop := 229 * 324 * c4 * 2
	poster := 297 * 420 * a3 * 2
	sheet := 210 * 297 * a4

	fmt.Printf("%.6f\n", (envelop+poster+sheet)*0.000001)
}
