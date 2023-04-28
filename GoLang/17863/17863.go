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

	if n%5550000 < 10000 && n/5550000 == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
