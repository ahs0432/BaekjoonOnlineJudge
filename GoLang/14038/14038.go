package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0

	for i := 0; i < 6; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)

		if tmp == "W" {
			sum++
		}
	}

	if sum >= 5 {
		fmt.Println("1")
	} else if sum >= 3 {
		fmt.Println("2")
	} else if sum >= 1 {
		fmt.Println("3")
	} else {
		fmt.Println("-1")
	}
}
