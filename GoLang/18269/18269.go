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

	var s, tmp string
	fmt.Fscanln(reader, &s)

	var check bool
	for k := 1; k <= n; k++ {
		check = true
		for i := 0; i <= n-k; i++ {
			tmp = s[i : i+k]
			for j := 0; j <= n-k; j++ {
				if i == j {
					continue
				}
				if tmp == s[j:j+k] {
					check = false
				}
			}
		}
		if check {
			fmt.Println(k)
			break
		}
	}
}
