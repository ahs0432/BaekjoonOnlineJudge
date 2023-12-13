package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k, tmp int
	fmt.Fscanln(reader, &n, &k)

	sum := 0
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp%2 == 0 {
			sum += tmp / 2
		} else {
			sum += tmp/2 + 1
		}
	}

	if n <= sum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
