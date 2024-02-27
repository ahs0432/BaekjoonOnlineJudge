package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	sum := int64(0)
	var tmp int64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		sum += int64(math.Pow(float64(tmp/10), float64(tmp%10)))
	}

	fmt.Println(sum)
}
