package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	if len(s) > 2 && s[0] == '"' && s[len(s)-1] == '"' {
		fmt.Println(s[1 : len(s)-1])
	} else {
		fmt.Println("CE")
	}
}
