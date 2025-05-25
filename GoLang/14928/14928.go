package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n big.Int
	fmt.Fscanln(reader, &n)

	fmt.Println(n.Rem(&n, big.NewInt(20000303)).String())
}
