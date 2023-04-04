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

	if (n/1000) >= 1 || (n/10) > 10 {
		fmt.Println(n/100 + n%100)
	} else {
		fmt.Println(n/10 + n%10)
	}
}
