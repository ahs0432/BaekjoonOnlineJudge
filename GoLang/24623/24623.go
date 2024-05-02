package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, a float64
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &a)

	fmt.Println(int(math.Round((180-a)/2 + a/2)))
}
