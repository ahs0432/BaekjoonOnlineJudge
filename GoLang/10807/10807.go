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

	t := map[int]int{}

	for i := 0; i < n; i++ {
		var tmp int
		if i != n-1 {
			fmt.Fscan(reader, &tmp)
		} else {
			fmt.Fscanln(reader, &tmp)
		}
		t[tmp]++
	}

	var v int
	fmt.Fscanln(reader, &v)
	fmt.Println(t[v])
}
