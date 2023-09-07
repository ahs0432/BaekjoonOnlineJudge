package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, i int
	fmt.Fscanln(reader, &a, &i)
	fmt.Println(a*(i-1) + 1)
}
