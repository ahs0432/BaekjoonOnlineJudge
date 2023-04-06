package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, t, d int
	fmt.Fscanln(reader, &s, &t, &d)
	fmt.Println(d / (s * 2) * t)
}
