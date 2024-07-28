package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, e int
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &e)

	prime := make([]bool, e+1)
	prime[0], prime[1] = true, true
	for i := 0; i*i <= e; i++ {
		if !prime[i] {
			for j := i * i; j <= e; j += i {
				prime[j] = true
			}
		}
	}

	min := 0
	sum := 0
	for ; s <= e; s++ {
		if !prime[s] {
			if min == 0 {
				min = s
			}
			sum += s
		}
	}

	if sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
}
