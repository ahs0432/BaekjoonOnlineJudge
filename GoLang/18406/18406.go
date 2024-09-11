package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	a := 0
	for i := 0; i < len(s)/2; i++ {
		a += int(s[i] - 48)
	}

	b := 0
	for i := len(s) / 2; i < len(s); i++ {
		b += int(s[i] - 48)
	}

	if a == b {
		fmt.Println("LUCKY")
	} else {
		fmt.Println("READY")
	}
}
