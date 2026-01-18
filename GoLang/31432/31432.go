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

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}

	if n == 1 {
		if d[0] == 0 {
			fmt.Println("YES")
			fmt.Println("0")
			return
		}
	}

	fmt.Println("YES")
	fmt.Println(d[n-1] * 111)
}
