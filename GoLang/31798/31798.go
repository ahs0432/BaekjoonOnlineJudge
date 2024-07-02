package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	if a == 0 {
		fmt.Println(c*c - b)
	} else if b == 0 {
		fmt.Println(c*c - a)
	} else if c == 0 {
		fmt.Println(int(math.Sqrt(float64(a + b))))
	}
}
