package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var b int
	fmt.Fscanln(reader, &b)

	kPa := b*5 - 400
	fmt.Println(kPa)

	if kPa > 100 {
		fmt.Println(-1)
	} else if kPa < 100 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
