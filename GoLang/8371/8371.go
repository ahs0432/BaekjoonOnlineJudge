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

	var t1 string
	var t2 string
	fmt.Fscanln(reader, &t1)
	fmt.Fscanln(reader, &t2)

	cnt := 0
	for i := 0; i < n; i++ {
		if t1[i] != t2[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
