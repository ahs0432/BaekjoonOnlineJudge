package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	apple := 0
	banana := 0

	for i := 3; i > 0; i-- {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		apple += (tmp * i)
	}

	for i := 3; i > 0; i-- {
		var tmp int
		fmt.Fscanln(reader, &tmp)
		banana += (tmp * i)
	}

	if apple == banana {
		fmt.Println("T")
	} else if apple > banana {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
