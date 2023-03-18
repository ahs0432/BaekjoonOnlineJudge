package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0

	for i := 0; i < 4; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		sum += tmp
	}

	fmt.Println(sum / 60)
	fmt.Println(sum % 60)
}
