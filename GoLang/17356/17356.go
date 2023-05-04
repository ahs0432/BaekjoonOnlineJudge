package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m float64
	fmt.Fscanln(reader, &n, &m)
	fmt.Printf("%.20f", 1/(1+math.Pow(10, (m-n)/400)))
}
