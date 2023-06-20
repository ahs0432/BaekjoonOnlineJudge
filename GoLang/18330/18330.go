package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &k)

	x := n
	if k+60 < n {
		x = k + 60
	}

	fmt.Println(x*1500 + (n-x)*3000)
}
