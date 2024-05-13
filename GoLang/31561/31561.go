package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscanln(reader, &n)

	var ans float64
	if n < 30 {
		ans = n / 2
	} else {
		ans = 15 + (n-30)*3/2
	}
	fmt.Printf("%.1f\n", ans)
}
