package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := make([]int, 6)
	for i := 0; i < 6; i++ {
		if i == 5 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	b := make([]int, 6)
	for i := 0; i < 6; i++ {
		if i == 5 {
			fmt.Fscanln(reader, &b[i])
		} else {
			fmt.Fscan(reader, &b[i])
		}
	}

	var answer int64 = 0
	for _, elementA := range a {
		for _, elementB := range b {
			if elementB < elementA {
				answer++
			}
		}
	}

	denominator := int64(len(a)) * int64(len(a))

	if answer == 0 {
		fmt.Println("0/1")
		return
	}

	common := gcd(answer, denominator)
	fmt.Printf("%d/%d\n", answer/common, denominator/common)
}
