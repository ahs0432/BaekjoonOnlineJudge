package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t, p float64
	fmt.Fscanln(reader, &t, &p)

	if p > 20 {
		fmt.Printf("%.6f\n", (p-20)*(t/(100-p))+20*2*(t/(100-p)))
	} else {
		fmt.Printf("%.6f\n", p*2*(t/((20-p)*2+80)))
	}
}
