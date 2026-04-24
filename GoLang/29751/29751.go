package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w, h float64
	fmt.Fscanln(reader, &w, &h)
	fmt.Printf("%.1f\n", w*h/2)
}
