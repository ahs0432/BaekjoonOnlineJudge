package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k float64
	fmt.Fscanln(reader, &k)

	var d1, d2 float64
	fmt.Fscanln(reader, &d1, &d2)

	fmt.Println(k*k - ((d1-d2)/2)*((d1-d2)/2))
}
