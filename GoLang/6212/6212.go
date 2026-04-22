package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(reader, &m, &n)
	digit := [10]int{}
	for i := m; i <= n; i++ {
		temp := i
		for temp > 0 {
			digit[temp%10]++
			temp /= 10
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", digit[i])
	}
	fmt.Println()
}
