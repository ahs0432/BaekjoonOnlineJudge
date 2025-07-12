package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p, k, tmp big.Int
	fmt.Fscanln(reader, &p, &k)

	check := true
	for i := big.NewInt(2); i.Cmp(&k) == -1; i.Add(i, big.NewInt(1)) {
		tmp.Mod(&p, i)
		if tmp.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("BAD", i.String())
			check = false
			break
		}
	}

	if check {
		fmt.Println("GOOD")
	}
}
