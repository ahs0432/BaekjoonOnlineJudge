package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d1, d2, x float64
	fmt.Fscanln(reader, &d1, &d2, &x)

	if d1 > d2 {
		d1, d2 = d2, d1
	}

	p := (100.0 - x) / x
	result := ((p + 1) / (p*float64(d2)/float64(d1) + 1)) * float64(d2)
	fmt.Printf("%.6f\n", result)
}
