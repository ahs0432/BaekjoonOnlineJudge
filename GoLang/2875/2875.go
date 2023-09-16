package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m, &k)

	for i := 0; i < k; i++ {
		if n/2 >= m {
			n -= 1
		} else {
			m -= 1
		}
	}

	if n/2 > m {
		fmt.Println(m)
	} else {
		fmt.Println(n / 2)
	}
}
