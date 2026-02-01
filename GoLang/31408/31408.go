package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x, maxCount int
	fmt.Fscanln(reader, &n)

	a := make([]int, 100001)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		a[x]++
		if a[x] > maxCount {
			maxCount = a[x]
		}
	}

	if maxCount <= (n+1)/2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
