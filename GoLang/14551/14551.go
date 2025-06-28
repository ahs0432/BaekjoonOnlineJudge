package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, tmp int
	fmt.Fscanln(reader, &n, &m)

	num := 1
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		if tmp == 0 {
			tmp = 1
		}
		num = num * tmp % m
	}
	fmt.Println(num % m)
}
