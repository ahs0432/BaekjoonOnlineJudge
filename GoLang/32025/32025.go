package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscanln(reader, &h)
	fmt.Fscanln(reader, &w)

	fmt.Println(min(h, w) * 100 / 2)
}
