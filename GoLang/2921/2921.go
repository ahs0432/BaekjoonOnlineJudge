package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(1.5 * float64(i*(i+1)))
	}
	fmt.Println(sum)
}
