package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, z, u, v, w int
	fmt.Fscanln(reader, &x, &y, &z)
	fmt.Fscanln(reader, &u, &v, &w)

	fmt.Println(u/100*x + v/50*y + w/20*z)
}
