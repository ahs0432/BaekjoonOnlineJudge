package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, f int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &f)

	n = n / 100 * 100
	ans := 0

	for (n+ans)%f != 0 && ans < 100 {
		ans++
	}
	fmt.Printf("%02d\n", ans)
}
