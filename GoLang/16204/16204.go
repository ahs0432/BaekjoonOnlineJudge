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

	sum := 0
	if m < k {
		sum = m
	} else {
		sum = k
	}

	if (n - m) < (n - k) {
		sum += (n - m)
	} else {
		sum += (n - k)
	}

	fmt.Println(sum)
}
