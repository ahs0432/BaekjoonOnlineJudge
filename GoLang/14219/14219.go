package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	if (n*m)%3 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
