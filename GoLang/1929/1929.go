package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(prime []bool) {
	prime[0] = true
	prime[1] = true

	for i := 2; i <= len(prime)/2; i++ {
		if !prime[i] {
			for j := 2; i*j < len(prime); j++ {
				prime[i*j] = true
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2 int
	fmt.Fscanln(reader, &n1, &n2)

	prime := make([]bool, n2+1)

	isPrime(prime)

	for ; n1 <= n2; n1++ {
		if !prime[n1] {
			fmt.Println(n1)
		}
	}
}
