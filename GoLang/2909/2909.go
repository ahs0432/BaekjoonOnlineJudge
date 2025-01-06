package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	var c, tmp float64
	fmt.Fscanln(reader, &c, &k)

	tmp = 1
	for i := 0; i < k; i++ {
		tmp *= 10
	}
	fmt.Println(int(math.Round((c+0.1)/tmp) * tmp))
}
