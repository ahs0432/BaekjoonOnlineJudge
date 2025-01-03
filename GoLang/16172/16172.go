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
	var k string

	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &k)

	var newS strings.Builder
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			newS.WriteString(string(c))
		}
	}

	if strings.Contains(newS.String(), k) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
