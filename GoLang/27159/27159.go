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

	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	r := 0
	f := t[0]
	p := t[0]

	for _, i := range t[1:] {
		if i != p+1 {
			r += f
			f = i
		}
		p = i
	}
	r += f
	fmt.Println(r)
}
