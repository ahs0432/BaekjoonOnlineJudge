package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var c, b float64
	fmt.Fscanln(reader, &c, &b)
	fmt.Printf("%.10f\n", c/b)
}
