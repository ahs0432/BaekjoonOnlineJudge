package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func sqrt(n uint64) *big.Int {
	ans := big.NewInt(1)
	for i := uint64(0); i < n; i++ {
		ans.Mul(ans, big.NewInt(3))
	}
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n uint64
	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		n -= 1
		fmt.Fprint(writer, "{ ")

		for i := uint64(0); n != 0; i++ {
			if n%2 == 1 {
				if n/2 == 0 {
					fmt.Fprint(writer, sqrt(i), " ")
				} else {
					fmt.Fprint(writer, sqrt(i), ", ")
				}
			}
			n /= 2
		}
		fmt.Fprintln(writer, "}")
	}

	writer.Flush()
}
