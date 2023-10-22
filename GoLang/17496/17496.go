package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, t, c, p int
	fmt.Fscanln(reader, &n, &t, &c, &p)
	fmt.Println((n - 1) / t * c * p)
}
