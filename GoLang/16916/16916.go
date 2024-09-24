package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, p string
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &p)

	if strings.Contains(s, p) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
