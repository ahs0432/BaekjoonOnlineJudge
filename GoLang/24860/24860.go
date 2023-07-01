package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var v1, v2, v3, j1, j2, j3, d1 int
	fmt.Fscanln(reader, &v1, &j1)
	fmt.Fscanln(reader, &v2, &j2)
	fmt.Fscanln(reader, &v3, &d1, &j3)

	fmt.Println(((v1 * j1) + (v2 * j2)) * (v3 * d1 * j3))
}
