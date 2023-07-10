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

	n = 25.0 + n*0.01
	if n < 100.00 {
		fmt.Println(100.00)
	} else if n > 2000.00 {
		fmt.Println(2000.00)
	} else {
		fmt.Printf("%.2f\n", n)
	}
}
