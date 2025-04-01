package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	fmt.Println(a + b - c)
	a, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	fmt.Println(a - c)
}
