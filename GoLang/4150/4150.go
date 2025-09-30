package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	f := []big.Int{*big.NewInt(0), *big.NewInt(1)}
	var sum *big.Int

	for i := 1; i <= n; i++ {
		sum = new(big.Int).Add(&f[0], &f[1])
		f[0].Set(&f[1])
		f[1].Set(sum)
	}
	fmt.Println(f[0].String())
}
