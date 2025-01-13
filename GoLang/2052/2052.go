package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscanln(reader, &n)

	ans := math.Pow(0.5, n)
	fmt.Println(strconv.FormatFloat(ans, 'f', int(n), 64))
}
