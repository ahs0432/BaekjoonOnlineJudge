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
	var st string
	fmt.Fscanln(reader, &st)
	count := 0

	for i := 0; i < n/2; i++ {
		if st[i] != st[n-1-i] {
			count++
		}
	}
	fmt.Println(count)
}
