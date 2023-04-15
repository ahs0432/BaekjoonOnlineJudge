package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, w, h, l int
	fmt.Fscanln(reader, &n, &w, &h, &l)

	if (w/l)*(h/l) > n {
		fmt.Println(n)
	} else {
		fmt.Println((w / l) * (h / l))
	}
}
