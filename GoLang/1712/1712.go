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

	if n2 > n3 || (n3-n2) == 0 {
		fmt.Println(-1)
	} else if n1 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(n1/(n3-n2) + 1)
	}
}
