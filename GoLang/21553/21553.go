package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, p string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &p)
	fmt.Println(p)
}
