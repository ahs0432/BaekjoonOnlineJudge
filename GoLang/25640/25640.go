package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var my string
	fmt.Fscanln(reader, &my)

	var n int
	fmt.Fscanln(reader, &n)

	some := 0
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		if tmp == my {
			some++
		}
	}

	fmt.Println(some)
}
