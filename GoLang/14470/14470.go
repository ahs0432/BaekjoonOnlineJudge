package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, e int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &e)

	sum := 0
	if a < 0 {
		sum += (int(math.Abs(float64(a))) * c)
		a = 0
	}

	if a == 0 {
		sum += d
	}

	sum += ((b - a) * e)

	fmt.Println(sum)
}
