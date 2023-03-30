package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	phi := 3.141592
	var d1, d2 float64

	fmt.Fscanln(reader, &d1)
	fmt.Fscanln(reader, &d2)

	sum := d1 * 2
	sum += (d2 * 2 * phi)
	fmt.Println(sum)
}
