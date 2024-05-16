package main

import (
	"bufio"
	"fmt"
	"os"
)

func prime(n int) []int {
	isPrime := make([]bool, n+1)
	isPrime[0] = true
	isPrime[1] = true

	for i := 2; i*i <= n; i++ {
		if !isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = true
			}
		}
	}

	primeNum := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			primeNum = append(primeNum, i)
		}
	}
	return primeNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	primeNum := prime(1000)
	table := make([]int, n)
	table[0] = 1
	table[1] = 2
	table[n-1] = 997

	for i := 2; i < n; i++ {
		if table[i] == 0 {
			table[i] = table[i-1] + 1
		}
	}

	fmt.Fprintln(writer, n)
	fmt.Fprint(writer, "1 2 ")
	for i := 2; i < n; i++ {
		fmt.Fprint(writer, primeNum[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
