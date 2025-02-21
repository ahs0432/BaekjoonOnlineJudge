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

	i := 1
	for ; i*(i+1)/2 < n; i++ {

	}
	b := n - (i-1)*i/2
	a := i + 1 - b
	fmt.Println(a, b)
}
