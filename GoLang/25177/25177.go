package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var a []int

	if n > m {
		a = make([]int, n)
	} else {
		a = make([]int, m)
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	res := 0
	var tmp int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &tmp)
		res = max(res, tmp-a[i])
	}
	fmt.Println(res)
}
