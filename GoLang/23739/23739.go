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

	sum, count := 0, 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if (tmp+1)/2 <= 30-sum {
			count++
		}

		sum += tmp
		if sum >= 30 {
			sum = 0
		}
	}
	fmt.Println(count)
}
