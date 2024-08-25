package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d, h, w float64
	fmt.Fscanln(reader, &d, &h, &w)
	fmt.Println(int(d/math.Sqrt(h*h+w*w)*h), int(d/math.Sqrt(h*h+w*w)*w))
}
