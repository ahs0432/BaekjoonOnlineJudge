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

	f := make([]int64, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &f[i])
		} else {
			fmt.Fscan(reader, &f[i])
		}
	}

	var c int64
	fmt.Fscanln(reader, &c)

	var tmp int64 = 0
	for i := 0; i < n; i++ {
		if f[i] > c {
			if f[i]%c == 0 {
				tmp += f[i] / c
			} else {
				tmp += f[i]/c + 1
			}
		} else if f[i] != 0 {
			tmp++
		}
	}

	fmt.Println(tmp * c)
}
