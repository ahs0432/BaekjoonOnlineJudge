package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, k, m int
	fmt.Fscanln(reader, &r, &k, &m)

	for i := m / k; i > 0 && r != 0; i-- {
		r >>= 1
	}

	fmt.Println(r)
}
