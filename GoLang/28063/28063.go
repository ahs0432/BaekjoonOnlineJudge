package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &x, &y)

	ans := 4
	if x == 1 {
		ans--
	}

	if x == n {
		ans--
	}

	if y == 1 {
		ans--
	}

	if y == n {
		ans--
	}
	fmt.Println(ans)
}
