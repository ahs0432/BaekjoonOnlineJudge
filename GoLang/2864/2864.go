package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a, &b)

	a6int, _ := strconv.Atoi(strings.ReplaceAll(a, "6", "5"))
	b6int, _ := strconv.Atoi(strings.ReplaceAll(b, "6", "5"))

	a5int, _ := strconv.Atoi(strings.ReplaceAll(a, "5", "6"))
	b5int, _ := strconv.Atoi(strings.ReplaceAll(b, "5", "6"))

	fmt.Println(a6int+b6int, a5int+b5int)
}
