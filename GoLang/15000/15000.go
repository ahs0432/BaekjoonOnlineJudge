package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	n = strings.ToUpper(n)
	fmt.Println(n)
}
