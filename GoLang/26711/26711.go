package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b big.Int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	fmt.Println(a.Add(&a, &b).String())
}
