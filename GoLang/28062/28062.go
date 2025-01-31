package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	sum := 0
	var tmp int
	odd := 1001
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp%2 == 1 && odd > tmp {
			odd = tmp
		}
		sum += tmp
	}

	if sum%2 == 1 {
		sum -= odd
	}
	fmt.Println(sum)
}
