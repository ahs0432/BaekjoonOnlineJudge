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

	var tmp big.Int
	var opr string
	var numTable []*big.Int
	var oprTable []string
	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Fscanln(reader, &opr)
		}
		fmt.Fscanln(reader, &tmp)

		if opr == "*" {
			tmp.Mul(numTable[len(numTable)-1], &tmp)
			numTable[len(numTable)-1] = new(big.Int).Set(&tmp)
		} else if opr == "/" {
			if tmp.Sign() == -1 && numTable[len(numTable)-1].Sign() == -1 {
				tmp.Div(numTable[len(numTable)-1].Abs(numTable[len(numTable)-1]), tmp.Abs(&tmp))
			} else if tmp.Sign() == -1 && numTable[len(numTable)-1].Sign() == 1 {
				tmp.Div(numTable[len(numTable)-1].Neg(numTable[len(numTable)-1]), tmp.Abs(&tmp))
			} else {
				tmp.Div(numTable[len(numTable)-1], &tmp)
			}
			numTable[len(numTable)-1] = new(big.Int).Set(&tmp)
		} else {
			numTable = append(numTable, new(big.Int).Set(&tmp))
			oprTable = append(oprTable, opr)
		}
	}

	ans := *numTable[0]
	for i := 1; i < len(numTable); i++ {
		if oprTable[i] == "+" {
			ans.Add(&ans, numTable[i])
		} else {
			ans.Sub(&ans, numTable[i])
		}
	}
	fmt.Println(ans.String())
}
