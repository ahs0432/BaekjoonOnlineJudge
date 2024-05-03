package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d string
	var n int
	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &n)

	var tmp string
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if d <= tmp {
			count++
		}
	}
	fmt.Println(count)
}
