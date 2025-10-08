package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s1, s2 string
	var b1, b2, ans big.Int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s1, &s2)
		b1 = *new(big.Int)
		b2 = *new(big.Int)

		b1.SetString(s1, 2)
		b2.SetString(s2, 2)

		ans.Add(&b1, &b2)
		fmt.Fprintln(writer, ans.Text(2))
	}
	writer.Flush()
}
