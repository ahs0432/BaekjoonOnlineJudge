package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	if strings.Contains(a, c) && strings.Contains(b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
