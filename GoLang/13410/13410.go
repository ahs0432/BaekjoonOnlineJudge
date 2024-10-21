package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	tmp := 0
	now := 0
	max := 0

	for i := 1; i <= k; i++ {
		now = n * i

		tmp = 0
		for now != 0 {
			tmp *= 10
			tmp += now % 10
			now /= 10
		}

		if tmp > max {
			max = tmp
		}
	}

	fmt.Println(max)
}
