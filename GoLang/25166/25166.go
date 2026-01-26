package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, m int
	fmt.Fscanln(reader, &s, &m)

	if s <= 1023 {
		fmt.Println("No thanks")
	} else if ((s - 1023) & m) == (s - 1023) {
		fmt.Println("Thanks")
	} else {
		fmt.Println("Impossible")
	}
}
