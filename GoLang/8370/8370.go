package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2, k1, k2 int
	fmt.Fscanln(reader, &n1, &k1, &n2, &k2)
	fmt.Println(n1*k1 + n2*k2)
}
