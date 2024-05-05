package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y, r int
	fmt.Fscanln(reader, &x, &y)
	fmt.Fscanln(reader, &r)

	fmt.Println(x-r, y+r)
	fmt.Println(x+r, y+r)
	fmt.Println(x+r, y-r)
	fmt.Println(x-r, y-r)
}
