package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, e int
	fmt.Fscanln(reader, &a, &b, &c, &d, &e)

	cnt := 0
	if a-1000 <= b {
		cnt++
	}

	if a-1000 <= c {
		cnt++
	}

	if a-1000 <= d {
		cnt++
	}

	if a-1000 <= e {
		cnt++
	}

	fmt.Println(cnt)
}
