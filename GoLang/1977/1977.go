package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, n int
	min, sum := 0, 0

	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &n)

	for i := 1; i*i <= n; i++ {
		if i*i >= m {
			sum += i * i
			if min == 0 {
				min = i * i
			}
		}
	}

	if sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
}
