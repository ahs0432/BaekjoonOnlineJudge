package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	se := strings.Count(s, "security")
	bd := strings.Count(s, "bigdata")

	if se > bd {
		fmt.Println("security!")
	} else if se < bd {
		fmt.Println("bigdata?")
	} else {
		fmt.Println("bigdata? security!")
	}
}
