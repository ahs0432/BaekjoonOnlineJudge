package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x int64
	fmt.Fscanln(reader, &n, &x)

	var before int64
	var first = true
	const mod int64 = 1000000007

	var ai, i int64
	for n >= 0 {
		fmt.Fscanln(reader, &ai, &i)
		if first {
			before = ai
			first = false
		} else {
			before = ((before * x) + ai) % mod
		}
		n--
	}

	fmt.Println(before)
}
