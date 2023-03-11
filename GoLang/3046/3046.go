package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, r1 int
	fmt.Fscanln(reader, &r1, &s)
	fmt.Println(s*2 - r1)
}
