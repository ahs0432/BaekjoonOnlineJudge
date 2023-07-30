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
	fmt.Printf("%.10f\n%.10f\n", 100/n, 100/(100-n))
}
