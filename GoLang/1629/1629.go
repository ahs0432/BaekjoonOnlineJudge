package main

import (
	"bufio"
	"fmt"
	"os"
)

func pow(a, b, c int64) int64 {
	if b == 1 {
		return a % c
	}

	tmp := pow(a, b/2, c)
	if b%2 == 0 {
		return tmp * tmp % c
	} else {
		return (a * ((tmp * tmp) % c) % c)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int64
	fmt.Fscanln(reader, &a, &b, &c)
	fmt.Println(pow(a, b, c))
}
