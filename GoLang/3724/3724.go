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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		matrix := make([][]int64, m)
		for r := 0; r < m; r++ {
			matrix[r] = make([]int64, n)
			for c := 0; c < n; c++ {
				fmt.Fscan(reader, &matrix[r][c])
			}
		}

		maxColIdx := 1
		maxProduct := big.NewInt(1)
		for r := 0; r < m; r++ {
			maxProduct.Mul(maxProduct, big.NewInt(matrix[r][0]))
		}

		for c := 1; c < n; c++ {
			currentProduct := big.NewInt(1)
			for r := 0; r < m; r++ {
				currentProduct.Mul(currentProduct, big.NewInt(matrix[r][c]))
			}

			if currentProduct.Cmp(maxProduct) >= 0 {
				maxProduct.Set(currentProduct)
				maxColIdx = c + 1
			}
		}

		fmt.Fprintln(writer, maxColIdx)
	}
	writer.Flush()
}
