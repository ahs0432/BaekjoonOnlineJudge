package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int64
	fmt.Fscanln(reader, &n, &m)
	fmt.Println(int64(math.Abs(float64(n - m))))
}
