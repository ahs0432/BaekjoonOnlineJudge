package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	var aCases []int
	var bCases []int

	limit := 1 << n
	for i := 0; i < limit; i++ {
		oneCnt := bits.OnesCount(uint(i))
		if oneCnt == a {
			aCases = append(aCases, i)
		}
		if oneCnt == b {
			bCases = append(bCases, i)
		}
	}

	maxAnswer := 0
	for _, valA := range aCases {
		for _, valB := range bCases {
			xorResult := valA ^ valB
			if xorResult > maxAnswer {
				maxAnswer = xorResult
			}
		}
	}

	fmt.Println(maxAnswer)
}
