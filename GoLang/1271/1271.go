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
	fmt.Fscanln(reader, &a, &b)

	div, divmod := new(big.Int).DivMod(&a, &b, new(big.Int))
	fmt.Println(div)
	fmt.Println(divmod)
}
