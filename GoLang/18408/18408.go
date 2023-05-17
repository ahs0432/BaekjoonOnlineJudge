package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2, n3 int
	fmt.Fscanln(reader, &n1, &n2, &n3)

	if n1+n2+n3 < 5 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
}
