package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var total int
	fmt.Fscanln(reader, &total)

	var tmp int
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &tmp)
		total -= tmp
	}

	fmt.Println(total)
}
